package main

import (
	"bufio"
	"fmt"
	"github.com/rkoesters/xdg/basedir"
	"github.com/rkoesters/xdg/desktop"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var pipe = ""
var runMode = true

func run(name string, path string) {
	if pipe != name {
		return
	}

	fmt.Println("Running: " + name)

	file, _ := os.Open(path)
	reader := bufio.NewReader(file)

	desktopFile, _ := desktop.New(reader)
	cmd := desktopFile.Exec
	tmpArray := strings.Fields(cmd)
	cmd = tmpArray[0]

	fmt.Println("Starting:" + cmd)

	exec.Command(cmd).Start()

	os.Exit(0)
}

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

	info, _ := os.Stdin.Stat()

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		runMode = false
	} else {
		scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))
		for scanner.Scan() {
			pipe = scanner.Text()
			fmt.Println(pipe)
		}
	}

	for _, v := range allDataDirs {
		var files []string

		err := filepath.Walk(v+"/applications", func(path string, info os.FileInfo, err error) error {
			if path[len(path)-8:] != ".desktop" {
				return nil
			}
			if runMode {
				run(name(path), path)
			} else {
				files = append(files, path)
			}

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
