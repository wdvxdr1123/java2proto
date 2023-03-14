package grammar

type Visitor interface {
	Visit(x JObject) (w Visitor)
}

// Walk 未完全完成，但是足够现在使用了
func Walk(v Visitor, x JObject) {
	if x == nil {
		return
	}
	// visit
	if v = v.Visit(x); v == nil {
		return
	}

	// visit children
	switch x := x.(type) {
	case *JProgramFile:
		walkList(v, x.Imports)
		walkList(v, x.TypeDecls)
	case *JClassDecl:
		walkList(v, x.Body)
	case *JClassBody:
		walkList(v, x.List)
	case *JMethodDecl:
		Walk(v, x.Block)
	case *JBlock:
		walkList(v, x.List)
	case *JSimpleStatement:
		Walk(v, x.Object)
	case *JAssignmentExpr:
		Walk(v, x.Left)
		Walk(v, x.Right)
	case *JVariableDecl:
		// todo
	}
}

func walkList(v Visitor, list []JObject) {
	for _, x := range list {
		Walk(v, x)
	}
}

type inspector func(JObject) bool

func (f inspector) Visit(obj JObject) Visitor {
	if f(obj) {
		return f
	}
	return nil
}

// Inspect traverses an AST in depth-first order: It starts by calling
// f(node); node must not be nil. If f returns true, Inspect invokes f
// recursively for each of the non-nil children of node, followed by a
// call of f(nil).
func Inspect(node JObject, f func(JObject) bool) {
	Walk(inspector(f), node)
}
