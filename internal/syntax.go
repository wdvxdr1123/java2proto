package internal

import (
	"fmt"

	"java2proto/internal/grammar"
)

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
