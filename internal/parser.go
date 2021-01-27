package internal

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"java2proto/internal/utils"
	"os"
	"strconv"
	"strings"
)

var (
	data  []byte
	index int
	err   error

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
{{range .ProtoItems}}  {{.Typename}} {{.Name}} = {{.Tag}};
{{end}}}
`

func format(name string) string {
	for _, prefix := range typePrefix {
		name = strings.TrimPrefix(name, prefix)
	}
	name = utils.SmallCamelCase(name)
	return name
}

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
				} else if sb.Len() == 0 || sb.String() == " " { // 空字符
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
	fail := map[string]string{}
L:
	for {
		token := nextToken()
		switch state {
		case 0:
			if token == "{" {
				state = 1
			} else if token == "}" {
				os.Exit(0)
			}
		case 1:
			switch token {
			case "class":
				messageName = nextToken()
			case "{":
				fields = []proto{}
				state = 2
			case "}":
				state = 1
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
					continue L
				}
				if peek(1) != "=" {
					if typeName != "static" {
						println("ignore", typeName, varName)
					}
					continue
				}
				enctype := convertTypeName(typeName)
				if !strings.HasPrefix(enctype, "repeated") {
					enctype = "optional " + enctype
				}
				for i := range fields {
					if fields[i].Name == varName {
						fields[i].Typename = enctype
						continue L
					}
				}
				fail[varName] = enctype
			case "}":
				if len(fields)+len(fail) == 0 { // ignore empty message
					continue
				}
				for k, v := range fail {
					lowerName := strings.ToLower(k)
					var similarity = 0
					var ind = 0
					for i := range fields {
						if fields[i].Typename != "" {
							continue
						}
						s := utils.Lccs(lowerName, utils.ToASCIILower(format(fields[i].Name)))
						if s > similarity {
							similarity = s
							ind = i
						}
					}
					println("auto match", fields[ind].Name, "->", k)
					fields[ind].Typename = v
					fields[ind].Name = func() string {
						if len(fields[ind].Name) > len(k) {
							return k
						}
						return fields[ind].Name
					}()
					delete(fail, k)
				}
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
		return "repeated " + getRepeatTypeName()
	}
	if prototype, ok := typenameMap[typename]; ok {
		return prototype
	}
	return MessagePrefix + typename
}

func getRepeatTypeName() string {
	var ret string
	_ = nextToken()                                            // =
	var class = nextToken()                                    // 类型段
	if strings.HasPrefix(class, "PBField.initRepeatMessage") { // repeat message
		ret = strings.TrimPrefix(class, "PBField.initRepeatMessage(")
		ret = strings.TrimSuffix(ret, ".class)")
		return ret
	} else { // PBField: repeat fiel
		ret = strings.TrimPrefix(class, "PBField.initRepeat(")
		ret = strings.TrimSuffix(ret, ".__repeatHelper__)")
		return convertTypeName(ret)
	}
}
