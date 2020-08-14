package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/rkoesters/xdg/basedir"
	"github.com/rkoesters/xdg/desktop"
)

var pipe = ""
var runMode = true

func run(name string, path string) {
	if pipe != name {
		return
	}

	terminal := "xterm"

	if len(os.Args) > 1 {
		terminal = os.Args[1]
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	desktopFile, err := desktop.New(reader)
	if err != nil {
		panic(err)
	}

	cmd := desktopFile.Exec
	cmdArray := strings.Fields(cmd)
	cmd = cmdArray[0]

	fmt.Println("Starting: " + cmd)
	if desktopFile.Terminal {
		exec.Command(terminal, "-e", cmd).Start()
	} else {
		exec.Command(cmd).Start()
	}

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
	if runMode {
		fmt.Println("Error: " + pipe + " not found")
		os.Exit(1)
	}
}
