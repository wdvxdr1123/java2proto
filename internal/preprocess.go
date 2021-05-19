package internal

import (
	"strings"
)

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
				return strings.TrimSpace(sb.String())
			}
		}
		sb.WriteByte(s.Data[s.Index])
		s.move(1)
	}
}
