package loader

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/wdvxdr1123/java2proto/internal/grammar"
	"github.com/wdvxdr1123/java2proto/internal/utils"
)

type Class struct {
	Name   string
	Outer  string
	Inners []*Class
	Types  map[string]string
	Tags   map[string]int
}

func NewClass() *Class {
	cls := &Class{
		Inners: make([]*Class, 0, 8),
		Types:  make(map[string]string),
		Tags:   make(map[string]int),
	}
	return cls
}

func (c *Class) walkClassDecl(decl *grammar.JClassDecl) {
	c.Outer, c.Name = cutClassName(decl.Name)
	for _, body := range decl.Body {
		switch body := body.(type) {
		case *grammar.JClassBody:
			c.walkClassBody(body)

		case *grammar.JBlock:
			c.walkBlock(body)

		default:
			panic(fmt.Sprintf("unknown body type %T", body))
		}
	}
}

func (c *Class) walkClassBody(body *grammar.JClassBody) {
loop:
	for _, decl := range body.List {
		switch decl := decl.(type) {
		case *grammar.JClassDecl:
			cls := NewClass()
			cls.walkClassDecl(decl)
			c.Inners = append(c.Inners, cls)

		case *grammar.JVariableDecl:
			if decl.Name == "__fieldMap__" {
				if decl.Init != nil {
					init := decl.Init.Expr.(*grammar.JMethodAccess)
					c.walkFieldMapInit(init)
				}
				continue
			}
			if decl.Init != nil && decl.Init.Expr != nil {
				if _, ok := decl.Init.Expr.(*grammar.JLiteral); ok {
					// todo(wdvxdr): parse enum & gen code.
					continue
				}
			}
			typ := decl.TypeSpec.Name.String()
			switch typ {
			case "PBRepeatField", "PBRepeatMessageField":
				var rptType string
				switch {
				case decl.Init != nil:
					rptType = parseRepeatFieldType(decl.Init.Expr)
				case len(decl.TypeSpec.TypeArgs) > 0:
					rptType = c.typeName(decl.TypeSpec.TypeArgs[0].TypeSpec.Name.String())
				default:
					continue loop
				}
				typ = "repeated " + c.typeName(rptType)

			default:
				typ = "optional " + c.typeName(typ)
			}
			c.Types[decl.Name] = typ

		case *grammar.JBlock:
			c.walkBlock(decl)

		case *grammar.JMethodDecl:
			if decl.Name == c.Name {
				c.walkConstructor(decl)
			}
		}
	}
}

func parseRepeatFieldType(expr grammar.JObject) string {
	switch expr := expr.(type) {
	case *grammar.JMethodAccess:
		switch expr := expr.ArgList[0].(type) {
		case *grammar.JReferenceType:
			return expr.Name.FirstType()
		case *grammar.JNameDotObject:
			return expr.Name.String()
		default:
			panic("unknown arg expr in repeated field init")
		}
		return expr.ArgList[0].(*grammar.JReferenceType).Name.FirstType()
	default:
		panic("unknown expr in repeated field init")
	}
}

func (c *Class) walkConstructor(decl *grammar.JMethodDecl) {
	for _, obj := range decl.Block.List {
		switch obj := obj.(type) {
		case *grammar.JSimpleStatement:
			if assign, ok := obj.Object.(*grammar.JAssignmentExpr); ok {
				left := asJObjectDotName(assign.Left)
				if left == nil {
					continue
				}

				if right := asJMethodAccess(assign.Right); right != nil {
					switch right.Method {
					case "initRepeat", "initRepeatMessage": // repeat message
						rtype := parseRepeatFieldType(right)
						c.Types[left.Name.LastType()] = "repeated " + c.typeName(rtype)
					}
				}
			}
		}
	}
}

func (c *Class) walkBlock(block *grammar.JBlock) {
	for _, stmt := range block.List {
		switch stmt := stmt.(type) {
		case *grammar.JLocalVariableDecl:
		// ignore
		case *grammar.JIfElseStmt:
			if stmt.IfBlock != nil {
				c.walkBlock(stmt.IfBlock.(*grammar.JBlock))
			}
			if stmt.ElseBlock != nil {
				c.walkBlock(stmt.ElseBlock.(*grammar.JBlock))
			}
		case *grammar.JSimpleStatement:
			if stmt, ok := stmt.Object.(*grammar.JAssignmentExpr); ok {
				left := asJReferenceType(stmt.Left)
				if left != nil && left.Name.String() == "__fieldMap__" {
					c.walkFieldMapInit(asJMethodAccess(stmt.Right))
				}
			}
		}
	}
}

func (c *Class) walkFieldMapInit(init *grammar.JMethodAccess) {
	tagArray := init.ArgList[0].(*grammar.JArrayAlloc).Init
	nameArray := init.ArgList[1].(*grammar.JArrayAlloc).Init
	for i := range tagArray {
		tag := tagArray[i].Expr.(*grammar.JLiteral).Text
		name, _ := strconv.Unquote(nameArray[i].Expr.(*grammar.JLiteral).Text)
		t, _ := strconv.Atoi(tag)
		c.Tags[name] = t >> 3
	}
}

var typePrefix = [...]string{
	"rpt_", "opt_", "msg_", "string_", "bool_", "int32_",
	"int64_", "uint32_", "uint64_", "sint32_",
	"sint64_", "bytes_", "str_", "float_",
	"double_", "fixed32_", "fixed64_", "sfixed32",
	"sfixed64", "str",
}

func format(name string) string {
	for _, prefix := range typePrefix {
		name = strings.TrimPrefix(name, prefix)
	}
	name = utils.SmallCamelCase(name)
	return name
}

func (c *Class) print(prefix string) string {
	buf := new(bytes.Buffer)
	write := func(format string, a ...interface{}) {
		buf.WriteString(prefix)
		fmt.Fprintf(buf, format, a...)
	}
	write("message %s {", c.Name)
	type item struct {
		Type string
		Name string
		ID   int
	}
	var items []item
	var failed []item
	for k := range c.Tags {
		itm := item{
			Type: c.Types[k],
			Name: k,
			ID:   c.Tags[k],
		}

		switch {
		case itm.Type == "":
			failed = append(failed, itm)
		default:
			delete(c.Types, k)
			items = append(items, itm)
		}
	}
	for _, itm := range failed {
		var matched string
		var match = -1
		for k := range c.Types {
			lccs := utils.Lccs(strings.ToLower(k), strings.ToLower(itm.Name))
			if lccs > match {
				matched = k
				match = lccs
			}
		}
		itm.Type = c.Types[matched]
		delete(c.Types, matched)
		items = append(items, itm)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].ID < items[j].ID
	})
	for i, itm := range items {
		if i == 0 {
			write("\n")
		}
		write("  %s %s = %d;\n", itm.Type, format(itm.Name), itm.ID)
	}
	for _, inner := range c.Inners {
		fmt.Fprintf(buf, "\n%s", inner.print(prefix+"  "))
	}
	write("}\n")
	return buf.String()
}

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

	// repeat field
	"String":          "string",
	"ByteStringMicro": "bytes",
	"Integer":         "uint32",
	"Long":            "uint64",
}

func (c *Class) typeName(typename string) string {
	if prototype, ok := typenameMap[typename]; ok {
		return prototype
	}
	outer, typ := cutClassName(typename)
	if outer == c.Outer || outer == "" {
		return typ
	}
	return outer + "." + typ
}
