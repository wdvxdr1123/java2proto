package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"java2proto/internal"
	"os"
	"path/filepath"
	"strings"
)

// preprocessCmd represents the preprocess command
var preprocessCmd = &cobra.Command{
	Use:   "preprocess",
	Short: "preprocess of the origin code",
	Long:  `预处理文本，提取常量保存至数据库中。`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := os.Getwd()
		getFilelist(dir)
		_ = internal.DB.Close()
	},
}

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "java") {
			d, _ := os.ReadFile(path)
			s := internal.Source{Data: d, Index: 0}
			s.Preprocess()
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func init() {
	rootCmd.AddCommand(preprocessCmd)
}
