package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/wdvxdr1123/java2proto/internal/loader"
	"github.com/wdvxdr1123/java2proto/internal/utils"
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

		for _, file := range args {
			if !utils.IsExist(file) {
				appPath, _ := os.Getwd()
				file = path.Join(appPath, file)
				if !utils.IsExist(file) {
					fmt.Println("文件路径不存在!")
					return
				}
			}

			loader.Parse(file)
		}
	},
}

func init() {
	rootCmd.AddCommand(protoCmd)
}
