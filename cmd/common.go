package cmd

import (
	"fmt"
	"os"
	"runtime"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
	}
}

func exitIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
func exitIfErrorWithDetail(err error, detail string) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("%s: %s", detail, err))
	os.Exit(1)
}

func getWorkDir() string {
	dir, err := os.Getwd()
	if err != nil {
		exitIfError(fmt.Errorf("failed to get the current working durectory: %w", err))
	}
	return dir
}
