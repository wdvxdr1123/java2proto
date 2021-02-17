package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"java2proto/internal"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var ch = make(chan string, 8)

// preprocessCmd represents the preprocess command
var preprocessCmd = &cobra.Command{
	Use:   "preprocess",
	Short: "preprocess of the origin code",
	Long:  `预处理文本，提取常量保存至数据库中。`,
	Run: func(cmd *cobra.Command, args []string) {
		wg := sync.WaitGroup{}
		dir, _ := os.Getwd()
		for i := 0; i < 8; i++ {
			wg.Add(1)
			go func() {
				for file := range ch {
					d, _ := os.ReadFile(file)
					s := internal.Source{Data: d, Index: 0}
					s.Preprocess()
				}
				wg.Done()
			}()
		}
		getFilelist(dir)
		wg.Wait()
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
			ch <- path
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	close(ch)
}

func init() {
	rootCmd.AddCommand(preprocessCmd)
}
