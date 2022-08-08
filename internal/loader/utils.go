package loader

import (
	"strings"

	"github.com/wdvxdr1123/java2proto/internal/grammar"
)

func asJMethodAccess(obj grammar.JObject) *grammar.JMethodAccess {
	r, _ := obj.(*grammar.JMethodAccess)
	return r
}

func asJReferenceType(obj grammar.JObject) *grammar.JReferenceType {
	r, _ := obj.(*grammar.JReferenceType)
	return r
}

func asJObjectDotName(obj grammar.JObject) *grammar.JObjectDotName {
	r, _ := obj.(*grammar.JObjectDotName)
	return r
}

func cutClassName(name string) (outer, inner string) {
	before, after, ok := strings.Cut(name, "$")
	if ok {
		return before, after
	}
	return "", before
}
