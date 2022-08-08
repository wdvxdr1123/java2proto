package cmd

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"

	"github.com/wdvxdr1123/java2proto/internal/loader"
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

func init() {
	rootCmd.AddCommand(protoCmd)
}
