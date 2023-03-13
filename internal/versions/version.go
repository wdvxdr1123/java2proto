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
	grammar.Walk(&methodVisitor{
		name: "WtloginHelper",
		v:    wtloginHelperVisitor{},
	}, prog)

	prog = loadFile(filepath.Join(file, "tools", "util.java"))
	grammar.Walk(sdkVisitor{}, prog)
}

type methodVisitor struct {
	name string
	v    grammar.Visitor
}

func (v *methodVisitor) Visit(x grammar.JObject) grammar.Visitor {
	if x, ok := x.(*grammar.JMethodDecl); ok {
		if x.Name == v.name {
			return v.v
		}
	}
	return v
}

type wtloginHelperVisitor struct{}

func (v wtloginHelperVisitor) Visit(x grammar.JObject) grammar.Visitor {
	if x, ok := x.(*grammar.JAssignmentExpr); ok {
		switch format(x.Left) {
		case "this.mMainSigMap":
			Version.MainSigMap = parseInt(x.Right)
		case "this.mSubSigMap":
			Version.SubSigMap = parseInt(x.Right)
		case "this.mMiscBitmap":
			Version.MiscBitmap = parseInt(x.Right)
		case "this.mOpenAppid":
			Version.OpenAppid = parseInt(x.Right)
		}
	}
	return v
}

type sdkVisitor struct{}

func (v sdkVisitor) Visit(x grammar.JObject) grammar.Visitor {
	switch x := x.(type) {
	case *grammar.JVariableDecl:
		switch x.Name {
		case "BUILD_TIME":
			Version.BuildTime = parseInt(x.Init.Expr)
		case "SDK_VERSION":
			Version.SDKVersion, _ = strconv.Unquote(format(x.Init.Expr))
		case "SSO_VERSION":
			Version.SSOVersion = parseInt(x.Init.Expr)
		}
	}
	return v
}

func DumpBeacon(file string) {
	prog := loadFile(file)
	grammar.Walk(beaconVisitor{}, prog)
}

type beaconVisitor struct{}

func (v beaconVisitor) Visit(x grammar.JObject) grammar.Visitor {
	switch x := x.(type) {
	case *grammar.JVariableDecl:
		switch x.Name {
		case "PUBLIC_MAIN_APP_KEY":
			Version.AppKey, _ = strconv.Unquote(format(x.Init.Expr))
		}
	}
	return v
}
