package cmd

import (
	"fmt"
	"os"
)

func exitIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

func getWorkDir() string {
	dir, err := os.Getwd()
	if err != nil {
		exitIfError(fmt.Errorf("failed to get the current working durectory: %w", err))
	}
	return dir
}
