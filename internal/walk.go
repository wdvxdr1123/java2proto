package internal

import (
	"fmt"
	"strconv"

	"github.com/wdvxdr1123/java2proto/internal/grammar"
	"github.com/wdvxdr1123/java2proto/internal/utils"
)

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
					rptType = decl.TypeSpec.TypeArgs[0].TypeSpec.Name.String()
				default:
					continue loop
				}
				typ = "repeated " + utils.ConvertTypeName(rptType)

			default:
				typ = "optional " + utils.ConvertTypeName(typ)
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
				left := assign.Left.(*grammar.JObjectDotName).Name.LastType()

				if right, ok := assign.Right.(*grammar.JMethodAccess); ok {
					switch right.Method {
					case "initRepeat", "initRepeatMessage": // repeat message
						rtype := parseRepeatFieldType(right)
						c.Types[left] = "repeated " + utils.ConvertTypeName(rtype)
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
	c.Name = decl.Name
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

func (c *Class) walkFieldMapInit(init *grammar.JMethodAccess) {
	tagArray := init.ArgList[0].(*grammar.JArrayAlloc).Init
	nameArray := init.ArgList[1].(*grammar.JArrayAlloc).Init
	for i := range tagArray {
		tag := tagArray[i].Expr.(*grammar.JLiteral).Text
		name := nameArray[i].Expr.(*grammar.JLiteral).Text
		name, _ = strconv.Unquote(name)
		t, _ := strconv.Atoi(tag)
		c.Tags[name] = t >> 3
	}
}
