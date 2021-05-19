package internal

import "java2proto/internal/grammar"

func Parse2(path string) {
	lexer := grammar.NewFileLexer(path, false)
	grammar.JulyParse(lexer)
	decls := lexer.JavaProgram().TypeDecls
	cls := NewClass()
	for _, d := range decls {
		switch decl := d.(type) {
		case *grammar.JClassDecl:
			cls.walkClassDecl(decl)
		}
	}
	println(cls.Name)
}
