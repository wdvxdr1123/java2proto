package internal

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"

	"java2proto/internal/utils"
)

var (
	PackageName   string
	MessagePrefix string
	WriteFile     string
)

type proto struct {
	Name     string
	Tag      int
	Typename string
}

var splitSymbol = [...]byte{
	' ',
	'\t',
	',',
	';',
	'\n',
	'\r',
}

var typePrefix = [...]string{
	"rpt_", "opt_", "msg_", "string_", "bool_", "int32_",
	"int64_", "uint32_", "uint64_", "sint32_",
	"sint64_", "bytes_", "str_", "float_",
	"double_", "fixed32_", "fixed64_", "sfixed32",
	"sfixed64", "str",
}

const messageTemplate = `
message {{.MessageName}} {
{{range .ProtoItems}}  {{.Typename}} {{.Name}} = {{.TagID}};
{{end}}}
`

func format(name string) string {
	for _, prefix := range typePrefix {
		name = strings.TrimPrefix(name, prefix)
	}
	name = utils.SmallCamelCase(name)
	return name
}

func Parse(fileName string) {
	pData, err := ioutil.ReadFile(fileName)
	defer func() {
		_ = recover()
	}()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if WriteFile != "false" {
		f, _ := os.OpenFile(fileName+".proto", os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_TRUNC, 0755)
		defer func() { _ = f.Close() }()
		os.Stdout = f
	}

	fmt.Println(`syntax = "proto2";`)
	if PackageName != "" {
		fmt.Printf("\noption go_package = \".;%v\";\n", PackageName)
	}
	s := Source{Data: pData, Index: 0}
	s.parse()
}

func (s *Source) parseFieldMap() (fields []proto) {
	peekUntil := func(b byte) {
		for s.Data[s.Index] != b {
			s.Index++
		}
		return
	}
	s.move(35) // = MessageMicro.initFieldMap(new int[]
	if s.peek(1) != "{" {
		peekUntil(';')
		s.move(1)
		return
	}
	s.move(1)
	st := s.Index
	peekUntil('}')
	tags := strings.Split(string(s.Data[st:s.Index]), ",")
	s.move(16) // }, new String[]{
	st = s.Index
	peekUntil('}')
	names := strings.Split(string(s.Data[st:s.Index]), ",")
	for i := range names {
		name := strings.Trim(names[i], "\" ")
		tagstr := strings.Trim(tags[i], "\" ")
		tag, _ := strconv.Atoi(tagstr)
		tag >>= 3
		fields = append(fields, proto{Name: name, Tag: tag})
	}
	peekUntil(';')
	s.move(1)
	return
}

func (s *Source) parse() {
	state, dep, write := 0, 0, 0
	s._import = map[string]string{}
	var messageName string
	var fields []proto
	saveProtoStruct := func() {
		tmpl, err := template.New("proto").Parse(messageTemplate)
		if err != nil {
			panic(err)
		}
		for i := range fields {
			fields[i].Name = format(fields[i].Name)
		}
		err = tmpl.Execute(os.Stdout, struct {
			MessageName string
			ProtoItems  []proto
		}{
			MessageName: MessagePrefix + messageName,
			ProtoItems:  fields,
		})
		if err != nil {
			panic(err)
		}
	}
L:
	for {
		token := s.nextToken()
		switch state {
		case 0:
			switch token {
			case "package":
				s._package = s.nextToken()
			case "import":
				importFile := s.nextToken()
				splited := strings.SplitN(importFile, ".", -1)
				s._import[splited[len(splited)-1]] = importFile
			case "{":
				state = 1
			case "}":
				os.Exit(0)
			}
		case 1:
			switch token {
			case "Class":
				messageName = s.nextToken()
			case "{":
				fields = []proto{}
				dep = 1
				state = 2
			case "}":
				state = 1
			default:
			}
		case 2:
			switch token {
			case "public", "static":
				typeName := s.nextToken()
				if typeName == "final" {
					typeName = s.nextToken()
				}
				varName := s.nextToken()
				if varName == "__fieldMap__" && s.nextToken() == "=" { // pb 元信息
					fields = s.parseFieldMap()
					write = 1
					continue L
				}
				if s.peek(1) != "=" {
					if typeName == "Class" {
						if write != 0 { // ignore empty message
							write = 0
							saveProtoStruct()
						}
						// println("嵌套", varName)
						messageName = varName
					} else if typeName != "static" { // 嵌套
						println("ignore", typeName, varName)
					}
					continue
				}
				enctype := s.convertTypeName(typeName)
				if !strings.HasPrefix(enctype, "repeated") {
					enctype = "optional " + enctype
				}
				for i := range fields {
					if fields[i].Name == varName {
						fields[i].Typename = enctype
						continue L
					}
				}
			case "{":
				dep++
			case "}":
				dep--
				if dep == 0 {
					if write != 0 { // ignore empty message
						write = 0
						saveProtoStruct()
					}
					state = 1
				}
			}
		default:
		}
	}
}

func (s *Source) convertTypeName(typename string) string {
	var typenameMap = map[string]string{
		"PBBoolField":     "bool",
		"PBBytesField":    "bytes",
		"PBDoubleField":   "double",
		"PBEnumField":     "uint32", // 不知道处理
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
	if strings.HasPrefix(typename, "PBRepeat") { // repeat 字段
		return "repeated " + s.getRepeatTypeName()
	}
	if prototype, ok := typenameMap[typename]; ok {
		return prototype
	}
	return MessagePrefix + typename
}

func (s *Source) getRepeatTypeName() string {
	var ret string
	_ = s.nextToken()                                          // =
	var class = s.nextToken()                                  // 类型段
	if strings.HasPrefix(class, "PBField.initRepeatMessage") { // repeat message
		ret = strings.TrimPrefix(class, "PBField.initRepeatMessage(")
		ret = strings.TrimSuffix(ret, ".Class)")
		return ret
	} else { // PBField: repeat fiel
		ret = strings.TrimPrefix(class, "PBField.initRepeat(")
		ret = strings.TrimSuffix(ret, ".__repeatHelper__)")
		return s.convertTypeName(ret)
	}
}
