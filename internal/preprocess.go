package internal

import (
	"container/list"
	"log"
	"os"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
)

var db *leveldb.DB

func init() {
	var err error
	db, err = leveldb.OpenFile(".java2proto", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Source struct {
	_package string
	_import  map[string]string
	Data     []byte
	Index    int
}

func (s *Source) peek(n int) string {
	return string(s.Data[s.Index : s.Index+n])
}

func (s *Source) move(n int) {
	s.Index = s.Index + n
	return
}

func (s *Source) nextToken() (token string) {
	sb := &strings.Builder{}
L:
	for {
		if s.Index >= len(s.Data) {
			os.Exit(0)
		}
		var now = s.Data[s.Index]
		for _, symbol := range splitSymbol {
			if now == symbol {
				s.move(1)
				if sb.String() == "/*" { // jadx 好像只生成这种注释
					sb.Reset()
					for {
						s.move(1)
						if s.Data[s.Index] == '*' && s.peek(2) == "*/" {
							s.move(2)
							return s.nextToken() // 递归处理
						}
					}
				} else if sb.Len() == 0 || strings.TrimSpace(sb.String()) == "" { // 空字符
					sb.Reset()
					continue L
				}
				return strings.Trim(sb.String(), " ")
			}
		}
		sb.WriteByte(s.Data[s.Index])
		s.move(1)
	}
}

func (s *Source) Preprocess() {
	s._import = map[string]string{}
	var (
		bracketCnt = 0
		stack      = list.New()
		nowClass   = ""
	)
	for {
		token := s.nextToken()
		switch token {
		case "package":
			s._package = s.nextToken()
			nowClass = s._package
		case "import":
			importFile := s.nextToken()
			splited := strings.SplitN(importFile, ".", -1)
			s._import[splited[len(splited)-1]] = importFile
		case "public":
			if s.peek(12) == "static final" {
				s.move(12)
				_ = s.nextToken()                         // type
				name := nowClass + "." + s.nextToken()    // var name
				_ = s.nextToken()                         // =
				val := strings.Trim(s.nextToken(), " \"") // value
				_ = db.Put([]byte(name), []byte(val), nil)
				//fmt.Println(name, "=", val)
			}
		case "class":
			className := s.nextToken()
			nowClass = nowClass + "." + className
			stack.PushBack(className)
		case "{":
			bracketCnt++
		case "}":
			bracketCnt--
			if bracketCnt < stack.Len() {
				back := stack.Back()
				nowClass = strings.TrimSuffix(nowClass, "."+back.Value.(string))
				stack.Remove(back)
			}
		}
	}
}
