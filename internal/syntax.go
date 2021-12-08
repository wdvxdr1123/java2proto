package internal

import (
	"fmt"

	"github.com/wdvxdr1123/java2proto/internal/grammar"
)

type Class struct {
	Name   string
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

func Parse(path string) {
	lexer := grammar.NewFileLexer(path, false)
	grammar.JulyParse(lexer)
	program := lexer.JavaProgram()
	if program.Pkg.Name.String() == "com.tencent.mobileqq.pb" {
		return
	}

	decls := program.TypeDecls
	cls := NewClass()
	for _, d := range decls {
		switch decl := d.(type) {
		case *grammar.JClassDecl:
			cls.walkClassDecl(decl)
		}
	}

	fmt.Println(`syntax = "proto2";`)
	fmt.Println()
	fmt.Println(cls.print(""))
}
