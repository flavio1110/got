/*
Copyright © 2023 Flavio Silva flavio1110@gmaill.com
*/
package cmd

import (
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/spf13/cobra"
)

// nbCmd represents the nb command.
var nbCmd = &cobra.Command{
	Use:   "nb",
	Short: "Creates a new branch",
	Long: `Creates a new branch with the specified name. For example:

got nb feature/123`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("missing branch name. e.g. got nb feature/123")
		}
		path := getWorkDir()

		r, err := git.PlainOpen(path)
		exitIfError(err)

		branchName := plumbing.NewBranchReferenceName(args[0])

		w, err := r.Worktree()
		exitIfError(err)

		err = w.Checkout(&git.CheckoutOptions{Branch: branchName, Keep: true, Create: true})
		exitIfError(err)
	},
}

func init() {
	rootCmd.AddCommand(nbCmd)
}
