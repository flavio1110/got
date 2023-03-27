/*
Copyright © 2023 Flavio Silva flavio1110@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var statCmd = &cobra.Command{
	Use:   "stat",
	Short: "Get the summarized status of the local repo.",
	Long: `Get the summarized status of the local repo.

Command: git status -s`,
	Run: func(cmd *cobra.Command, args []string) {
		path := getWorkDir()

		repo, err := git.PlainOpen(path)
		exitIfError(err)

		wt, err := repo.Worktree()
		exitIfError(err)

		status, err := wt.Status()
		exitIfError(err)

		if status.IsClean() {
			fmt.Printf("%s All clean! \U0001FAE7\n", Blue)
			return
		}

		for path, change := range status {

			switch change.Staging {
			case git.Renamed:
				fmt.Printf("%s%s -> %s\n", Purple, path, change.Extra)
			default:
				fmt.Printf("%s%c%s%c%s - %s\n", Green, change.Staging, Red, change.Worktree, White, path)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(statCmd)
}
