package loader

import (
	"fmt"
	"os"

	"github.com/wdvxdr1123/java2proto/interna/loader"
	"github.com/wdvxdr1123/java2proto/internal/grammar"
)

type Package struct {
	Path    string
	Name    string
	Classes []*loader.Class
}

func LoadPackage(path string) (*Package, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			continue
		}
		info, _ := dir.Info()
		_ = info
		println(dir.Name())
	}
	return nil, nil
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
