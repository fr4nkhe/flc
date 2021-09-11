package main

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()
	err := filepath.Walk(pwd, func(p string, info os.FileInfo, err error) error {
		if path.Ext(p) == GoFilesSuffix {
			_, err := ReadContent(info.Name(), p)
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
