package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/watergist/file-engine/list"
	"github.com/watergist/file-engine/reader"
	"log"
	"os"
	"path"
	"strings"
)

func main() {
	var dir = GetFlags()

	tree := list.Read(&([]string{*dir}), -1)
	tree.SkipHiddenDirs = true
	err := tree.UpdateFiles()
	if err != nil {
		log.Fatal(err)
	}
	for _, foundFile := range tree.FilePaths {
		if strings.HasSuffix(foundFile, ".html") && !strings.Contains(foundFile, "css.html") {
			err := reader.CopyContent(path.Join("conv", foundFile),
				path.Join(*dir, "css.html"),
				foundFile)
			if errors.Is(err, os.ErrNotExist) {
				fmt.Println(err)
			}
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func GetFlags() *string {
	dir := flag.String("dir", "envoy", "Directory which needs to be modified")
	flag.Parse()
	return dir
}
