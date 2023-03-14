package cmd

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"

	"github.com/wdvxdr1123/java2proto/internal/loader"
	"github.com/wdvxdr1123/java2proto/internal/versions"
)

// protoCmd represents the proto command
var protoCmd = &cobra.Command{
	Use:   "proto",
	Short: "from java file generate proto file",
	Long:  `通过腾讯java文件导出proto文件`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("请输入你要转换的文件路径")
			return
		}
		var wg sync.WaitGroup
		filepath.Walk(args[0], func(path string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				wg.Add(1)
				go func() {
					defer wg.Done()
					pkg, err := loader.LoadPackage(path)
					if err != nil {
						return
					}
					pkg.Dump(args[1])
				}()
			}
			return nil
		})
		wg.Wait()
	},
}

// dump represents the dump command
var dump = &cobra.Command{
	Use:   "dump [output folder]",
	Short: "dump version info",
	Long:  `dump version info`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Help()
			return
		}

		versions.DumpWtloginSDK("oicq/wlogin_sdk")
		versions.DumpBeacon("./com/tencent/mobileqq/statistics/QQBeaconReport.java")
		versions.DumpAppSetting("./com/tencent/common/config/AppSetting.java")
		versions.FixUp()

		folder := args[0]

		output, _ := json.MarshalIndent(&versions.APhone, "", "  ")
		os.WriteFile(filepath.Join(folder, "android_phone.json"), output, 0o644)

		output, _ = json.MarshalIndent(&versions.APad, "", "  ")
		os.WriteFile(filepath.Join(folder, "android_pad.json"), output, 0o644)
	},
}

func init() {
	rootCmd.AddCommand(protoCmd)
	rootCmd.AddCommand(dump)
}
