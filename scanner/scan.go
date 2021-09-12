package scanner

import (
	"github.com/fr4nkhe/flc/parser"
	"log"
	"os"
	"path"
	"path/filepath"
)

func Scan() {
	pwd, _ := os.Getwd()
	err := filepath.Walk(pwd, func(p string, info os.FileInfo, err error) error {
		if path.Ext(p) == parser.GoFilesSuffix {
			_, err := parser.ReadContent(p)
			if err != nil {
				log.Fatalf("read file content error:%s", err)
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("read path error:%s", err)
	}
}
