package versions

import (
	"bytes"
	"encoding/hex"
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
	case *grammar.JReferenceType:
		return x.Name.String()
	}
	return ""
}

func parseInt(x grammar.JObject) uint64 {
	switch x := x.(type) {
	case *grammar.JLiteral:
		t := strings.Trim(x.Text, "L")
		i, _ := strconv.ParseUint(t, 10, 64)
		return i
	}
	return 0
}

func FixUp() {
	APhone.ApkId = "com.tencent.mobileqq"
	APhone.SubAppId = APhone.AppId
	APhone.ApkSign = hex.EncodeToString([]byte{0xA6, 0xB7, 0x45, 0xBF, 0x24, 0xA2, 0xC2, 0x77, 0x52, 0x77, 0x16, 0xF6, 0xF3, 0x6E, 0xB6, 0x8D})
	APhone.ProtocolType = AndroidPhone

	// copy from aphone
	id := APad.AppId
	APad = APhone
	APad.AppId = id
	APad.SubAppId = id
	APad.ProtocolType = AndroidPad
}
