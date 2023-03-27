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

// nbCmd represents the nb command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
