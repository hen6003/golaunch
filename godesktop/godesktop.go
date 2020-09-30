package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rkoesters/xdg/basedir"
	"github.com/rkoesters/xdg/desktop"
)

func name(path string) string {
	file, _ := os.Open(path)
	reader := bufio.NewReader(file)

	desktopFile, _ := desktop.New(reader)
	name := desktopFile.Name

	return name
}

func main() {
	allDataDirs := []string{basedir.DataHome}

	for _, v := range basedir.DataDirs {
		allDataDirs = append(allDataDirs, v)
	}

	for _, v := range allDataDirs {
		var files []string

		err := filepath.Walk(v+"/applications", func(path string, info os.FileInfo, err error) error {
			if path[len(path)-8:] != ".desktop" {
				return nil
			}

			files = append(files, path)

			return nil
		})
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			fmt.Println(name(file))
		}
	}
}
