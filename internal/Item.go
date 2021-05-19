package internal

import (
	"fmt"
	"strconv"

	"java2proto/internal/grammar"
	"java2proto/internal/utils"
)

type Class struct {
	Name   string
	Inners []*Class
	Types  map[string]string
	Tags   map[string]int
}

func NewClass() *Class {
	return &Class{
		Inners: []*Class{},
		Types:  map[string]string{},
		Tags:   map[string]int{},
	}
}

func (c *Class) walkClassBody(body *grammar.JClassBody) {
	for _, decl := range body.List {
		switch decl := decl.(type) {
		case *grammar.JClassDecl:
			cls := NewClass()
			cls.Name = decl.Name
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
			typ := decl.TypeSpec.Name.String()
			switch typ {
			case "PBRepeatField", "PBRepeatMessageField":
				var rptType string
				if len(decl.TypeSpec.TypeArgs) > 0 {
					rptType = decl.TypeSpec.TypeArgs[0].TypeSpec.Name.String()
				} else {
					rptType = parseRepeatFieldType(decl.Init)
				}
				typ = "repeated " + utils.ConvertTypeName(rptType)
			default:
				typ = utils.ConvertTypeName(typ)
			}
			c.Types[decl.Name] = typ

		case *grammar.JBlock:
			c.walkBlock(decl)
		}
	}
}

func (c *Class) walkBlock(block *grammar.JBlock) {
	for _, stmt := range block.List {
		switch stmt := stmt.(type) {
		case *grammar.JLocalVariableDecl:
		// ignore
		case grammar.JSimpleStatement:
			if stmt, ok := stmt.Object.(*grammar.JAssignmentExpr); ok {
				left := stmt.Left.(*grammar.JReferenceType)
				right := stmt.Right.(*grammar.JMethodAccess)
				if left.Name.String() != "__fieldMap__" {
					continue
				}
				c.walkFieldMapInit(right)
			}
		}
	}
}

func (c *Class) walkClassDecl(decl *grammar.JClassDecl) {
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

func parseRepeatFieldType(init *grammar.JVariableInit) string {
	switch expr := init.Expr.(type) {
	case grammar.JMethodAccess:
		return expr.ArgList[0].(*grammar.JReferenceType).Name.FirstType()
	default:
		panic("unknown expr in repeated field init")
	}
}

func (c *Class) walkFieldMapInit(init *grammar.JMethodAccess) {
	tagArray := init.ArgList[0].(*grammar.JArrayAlloc).Init
	nameArray := init.ArgList[1].(*grammar.JArrayAlloc).Init
	for i := range tagArray {
		tag := tagArray[i].Expr.(*grammar.JLiteral).Text
		name := nameArray[i].Expr.(*grammar.JLiteral).Text
		t, _ := strconv.Atoi(tag)
		c.Tags[name] = t
	}
}
