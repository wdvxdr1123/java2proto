package versions

import (
	"strconv"
	"strings"

	"github.com/wdvxdr1123/java2proto/internal/grammar"
)

type ProtocolType int

const (
	Unset        ProtocolType = iota
	AndroidPhone              // Android Phone
	AndroidWatch              // Android Watch
	MacOS                     // MacOS
	QiDian                    // 企点
	IPad                      // iPad
	AndroidPad                // Android Pad
)

type AppVersion struct {
	ApkId           string       `json:"apk_id"`
	AppId           uint64       `json:"app_id"`
	SubAppId        uint64       `json:"sub_app_id"`
	AppKey          string       `json:"app_key"`
	SortVersionName string       `json:"sort_version_name"`
	BuildTime       uint64       `json:"build_time"`
	ApkSign         string       `json:"apk_sign"` // hex encoded
	SdkVersion      string       `json:"sdk_version"`
	SSOVersion      uint64       `json:"sso_version"`
	MiscBitmap      uint64       `json:"misc_bitmap"`
	MainSigMap      uint64       `json:"main_sig_map"`
	SubSigmap       uint64       `json:"sub_sig_map"`
	ProtocolType    ProtocolType `json:"protocol_type"`
}

var (
	APhone AppVersion
	APad   AppVersion
)

func DumpAppSetting(file string) {
	prog := loadFile(file)
	grammar.Inspect(prog, func(x grammar.JObject) bool {
		switch x := x.(type) {
		case *grammar.JAssignmentExpr:
			left := format(x.Left)
			// TODO: 这样可靠吗？
			switch {
			case strings.HasSuffix(left, "f"):
				APhone.AppId = parseInt(x.Right)
			case strings.HasSuffix(left, "g"):
				APad.AppId = parseInt(x.Right)
			case strings.HasSuffix(left, "m"):
				APhone.SortVersionName, _ = strconv.Unquote(format(x.Right))
			}
			return false
		case *grammar.JMethodDecl:
			return false
		}
		return true
	})
}
