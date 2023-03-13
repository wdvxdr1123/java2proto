package versions

import (
	"bytes"
	"os"
	"strconv"
	"strings"

	"github.com/wdvxdr1123/java2proto/internal/grammar"
)

func loadFile(file string) *grammar.JProgramFile {
	readFile, err := os.ReadFile(file)
	if err != nil {
		return nil
	}
	// parser无法识别，替换以保证parse成功
	readFile = bytes.ReplaceAll(readFile, []byte("<>"), nil)
	readFile = bytes.ReplaceAll(readFile, []byte("??"), []byte("int"))
	lexer := grammar.NewLexer(bytes.NewReader(readFile), false)
	grammar.JulyParse(lexer)
	return lexer.JavaProgram()
}

func format(x grammar.JObject) string {
	switch x := x.(type) {
	case *grammar.JKeyword:
		return x.Name
	case *grammar.JObjectDotName:
		return format(x.Obj) + "." + x.Name.String()
	case *grammar.JLiteral:
		return x.Text
	}
	return ""
}

func parseInt(x grammar.JObject) int64 {
	switch x := x.(type) {
	case *grammar.JLiteral:
		t := strings.Trim(x.Text, "L")
		i, _ := strconv.ParseInt(t, 10, 64)
		return i
	}
	return 0
}
