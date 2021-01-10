package internal

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var (
	data  []byte
	index int
	err   error

	PackageName string
)

type proto struct {
	Name     string
	Tag      int
	Typename string
}

var splitSymbol = [...]byte{
	' ',
	',',
	';',
	'\n',
}

const messageTemplate = `
message {{.MessageName}} {
{{range .ProtoItems}}  {{.Typename}} {{.Name}} = {{.Tag}};
{{end}}}
`

func peek(n int) string {
	return string(data[index : index+n])
}

func move(n int) {
	index += n
	return
}

func nextToken() (token string) {
	sb := &strings.Builder{}
	for {
		if index >= len(data) {
			os.Exit(0)
		}
		for _, symbol := range splitSymbol {
			if data[index] == symbol {
				move(1)
				if sb.String() == "/*" { // jadx 好像只生成这种注释
					sb.Reset()
					for {
						move(1)
						if data[index] == '*' && peek(2) == "*/" {
							move(2)
							return nextToken() // 递归处理
						}
					}
				} else if sb.Len() == 0 || sb.String() == " " {
					sb.Reset()
					continue
				}
				return strings.Trim(sb.String(), " ")
			}
		}
		sb.WriteByte(data[index])
		move(1)
	}
}

func Parse(fileName string) {
	data, err = ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(`syntax = "proto2";`)
	if PackageName != "" {
		fmt.Printf("\noption go_package = \".;%v\";\n", PackageName)
	}
	parse()
}

func parseFieldMap() (fields []proto) {
	peekUntil := func(b byte) {
		for data[index] != b {
			index++
		}
		return
	}
	move(37) // = MessageMicro.initFieldMap(new int[]
	if peek(1) != "{" {
		peekUntil(';')
		move(1)
		return
	}
	move(1)
	st := index
	peekUntil('}')
	tags := strings.SplitN(string(data[st:index]), ",", -1)
	move(16) // }, new String[]{
	st = index
	peekUntil('}')
	names := strings.SplitN(string(data[st:index]), ",", -1)
	for i := range names {
		name := strings.Trim(names[i], "\" ")
		tagstr := strings.Trim(tags[i], "\" ")
		tag, _ := strconv.Atoi(tagstr)
		tag >>= 3
		fields = append(fields, proto{Name: name, Tag: tag})
	}
	peekUntil(';')
	move(1)
	return
}

func parse() {
	state := 0
	var messageName string
	var fields []proto
	saveProtoStruct := func() {
		tmpl, err := template.New("proto").Parse(messageTemplate)
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(os.Stdout, struct {
			MessageName string
			ProtoItems  []proto
		}{
			MessageName: messageName,
			ProtoItems:  fields,
		})
		if err != nil {
			panic(err)
		}
	}

	for {
		token := nextToken()
		switch state {
		case 0:
			if token == "{" {
				state = 1
			}
		case 1:
			switch token {
			case "class":
				messageName = nextToken()
				println(messageName)
			case "{":
				state = 2
			default:
			}
		case 2:
			switch token {
			case "public", "static":
				typeName := nextToken()
				if typeName == "final" {
					typeName = nextToken()
				}
				varName := nextToken()
				if varName == "__fieldMap__" { // pb 元信息
					fields = parseFieldMap()
				}
				for i := range fields {
					if fields[i].Name == varName {
						enctype := convertTypeName(typeName)
						if !strings.HasPrefix(enctype, "repeated") {
							enctype = "optional " + enctype
						}
						fields[i].Typename = enctype
					}
				}
			case "}":
				saveProtoStruct()
				state = 1
			}
		default:
		}
	}
}

func convertTypeName(typename string) string {
	var typenameMap = map[string]string{
		"PBBoolField":     "bool",
		"PBBytesField":    "bytes",
		"PBDoubleField":   "double",
		"PBEnumField":     "int32", // 不知道处理
		"PBFixed32Field":  "fixed32",
		"PBFixed64Field":  "fixed64",
		"PBFloatField":    "float",
		"PBInt32Field":    "int32",
		"PBInt64Field":    "int64",
		"PBSFixed32Field": "sfixed32",
		"PBSFixed64Field": "sfixed64",
		"PBSInt32Field":   "sint32",
		"PBSInt64Field":   "sint64",
		"PBStringField":   "string",
		"PBUInt32Field":   "uint32",
		"PBUInt64Field":   "uint64",
	}
	if prototype, ok := typenameMap[typename]; ok {
		return prototype
	}
	typename = strings.TrimPrefix(typename, "PBRepeat")
	typename = strings.TrimPrefix(typename, "PBRepeatMessageField")
	if strings.HasPrefix(typename, "<") {
		return "repeated " + convertTypeName(strings.Trim(typename, "<>"))
	}
	return typename
}
