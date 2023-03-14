package versions

import (
	"path/filepath"
	"strconv"

	"github.com/wdvxdr1123/java2proto/internal/grammar"
)

var Version struct {
	MainSigMap int64  `json:"main_sig_map"`
	MiscBitmap int64  `json:"misc_bitmap"`
	OpenAppid  int64  `json:"open_appid"`
	SubSigMap  int64  `json:"sub_sig_map"`
	SSOVersion int64  `json:"sso_version"`
	BuildTime  int64  `json:"build_time"`
	SDKVersion string `json:"sdk_version"`
	AppKey     string `json:"app_key"`
}

func DumpWtloginSDK(file string) {
	prog := loadFile(filepath.Join(file, "request", "WtloginHelper.java"))

	grammar.Inspect(prog, func(x grammar.JObject) bool {
		if x, ok := x.(*grammar.JMethodDecl); ok {
			if x.Name == "WtloginHelper" {
				walkWtloginHelper(x)
			}
			return false
		}
		return true
	})

	prog = loadFile(filepath.Join(file, "tools", "util.java"))
	grammar.Inspect(prog, func(x grammar.JObject) bool {
		switch x := x.(type) {
		case *grammar.JVariableDecl:
			switch x.Name {
			case "BUILD_TIME":
				APhone.BuildTime = parseInt(x.Init.Expr)
			case "SDK_VERSION":
				APhone.SdkVersion, _ = strconv.Unquote(format(x.Init.Expr))
			case "SSO_VERSION":
				APhone.SSOVersion = parseInt(x.Init.Expr)
			}
		}
		return true
	})
}

func walkWtloginHelper(x grammar.JObject) {
	grammar.Inspect(x, func(x grammar.JObject) bool {
		if x, ok := x.(*grammar.JAssignmentExpr); ok {
			switch format(x.Left) {
			case "this.mMainSigMap":
				APhone.MainSigMap = parseInt(x.Right)
			case "this.mSubSigMap":
				APhone.SubSigmap = parseInt(x.Right)
			case "this.mMiscBitmap":
				APhone.MiscBitmap = parseInt(x.Right)
			case "this.mOpenAppid":
				// APhone.OpenAppid = parseInt(x.Right)
			}
		}
		return true
	})
}

func DumpBeacon(file string) {
	prog := loadFile(file)
	grammar.Inspect(prog, func(x grammar.JObject) bool {
		switch x := x.(type) {
		case *grammar.JVariableDecl:
			switch x.Name {
			case "PUBLIC_MAIN_APP_KEY":
				APhone.AppKey, _ = strconv.Unquote(format(x.Init.Expr))
			}
		}
		return true
	})
}
