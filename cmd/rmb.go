/*
Copyright © 2023 Flavio Silva flavio1110@gmaill.com
*/
package cmd

import (
	"fmt"
	"io"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var rmbCmd = &cobra.Command{
	Use:   "rmb",
	Short: "Remove all local branches except the current one, main, and master",
	Long: `Remove all local branches except the current one, main, and master

	e.g. alias gbr="git branch | grep -v "main" | xargs git branch -D"`,
	Run: func(cmd *cobra.Command, args []string) {
		path := getWorkDir()

		repo, err := git.PlainOpen(path)
		exitIfError(err)

		branches, err := repo.Branches()
		exitIfError(err)

		skipList := [...]string{"main", "master"}
		for {
			branch, err := branches.Next()

			if err == io.EOF {
				break
			}

			exitIfErrorWithDetail(err, "Next branch")

			if branch == nil {
				break
			}

			name := branch.Name().Short()
			shouldSkip := false
			for _, skipItem := range skipList {
				if skipItem == name {
					shouldSkip = true
					break
				}
			}
			if shouldSkip {
				fmt.Printf("Skipped %q\n", name)
				continue
			}

			err = repo.Storer.RemoveReference(branch.Name())
			exitIfErrorWithDetail(err, fmt.Sprintf("delete branch %q", name))
			fmt.Printf("%sDeleted %q%s\n", Red, name, White)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmbCmd)
}
