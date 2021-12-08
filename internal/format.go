package internal

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"java2proto/internal/utils"
)

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
