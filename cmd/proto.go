package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"java2proto/internal"
	"java2proto/internal/utils"
	"os"
	"path"
)

// protoCmd represents the proto command
var protoCmd = &cobra.Command{
	Use:   "proto",
	Short: "from java file generate proto file",
	Long:  `通过腾讯java文件导出proto文件`,
	Run: func(cmd *cobra.Command, args []string) {
		if args == nil || len(args) == 0 {
			fmt.Println("请输入你要转换的文件路径")
			return
		}

		internal.PackageName = cmd.Flag("package").Value.String()
		internal.MessagePrefix = cmd.Flag("prefix").Value.String()

		file := args[0]
		if !utils.IsExist(file) {
			appPath, _ := os.Getwd()
			file = path.Join(appPath, file)
			if !utils.IsExist(file) {
				fmt.Println("文件路径不存在!")
				return
			}
		}

		internal.Parse(file)
	},
}

func init() {
	rootCmd.AddCommand(protoCmd)
	protoCmd.Flags().String("package", "", "go_package name")

	protoCmd.Flags().String("prefix", "", "message prefix name")
}
