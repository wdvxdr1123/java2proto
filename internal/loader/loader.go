package loader

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/wdvxdr1123/java2proto/internal/grammar"
)

type Package struct {
	Path    string
	Name    string
	Classes map[string][]*Class
}

func LoadPackage(path string) (*Package, error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	pkg := &Package{
		Path:    path,
		Classes: make(map[string][]*Class),
	}
	for _, f := range dir {
		if f.IsDir() {
			continue
		}
		pkg.load(f.Name())
	}
	return pkg, nil
}

func (pkg *Package) load(file string) {
	if path.Ext(file) != ".java" {
		return
	}
	lexer := grammar.NewFileLexer(path.Join(pkg.Path, file), false)
	grammar.JulyParse(lexer)
	prog := lexer.JavaProgram()
	if prog == nil {
		println("error at:", path.Join(pkg.Path, file))
		return
	}
	if prog.Pkg.Name.String() == "com.tencent.mobileqq.pb" {
		return
	}

	var ok bool
	for _, imp := range prog.Imports {
		ok = ok || imp.(*grammar.JImportStmt).Name.PackageString() == "com.tencent.mobileqq.pb"
	}
	if !ok {
		return
	}

	pkg.decls(prog.TypeDecls)
}

func (pkg *Package) decls(list []grammar.JObject) {
	for _, d := range list {
		switch decl := d.(type) {
		case *grammar.JClassDecl:
			outer, inner := cutClassName(decl.Name)
			cls := NewClass()
			cls.walkClassDecl(decl)
			if outer == "" {
				outer = inner
			}
			pkg.Classes[outer] = append(pkg.Classes[outer], cls)
		}
	}
}

func (pkg *Package) Dump(p string) {
	buf := new(bytes.Buffer)
	p = path.Join(p, pkg.Path)
	if len(pkg.Classes) != 0 {
		os.MkdirAll(p, 0750)
	}
	for name, classes := range pkg.Classes {
		buf.Reset()
		fmt.Fprintln(buf, `syntax = "proto2";`)
		for i, class := range classes {
			if i == 0 {
				fmt.Fprintln(buf)
			}
			fmt.Fprintln(buf, class.print(""))
		}
		_ = os.WriteFile(path.Join(p, name+".proto"), buf.Bytes(), 0o644)
	}
}
