/*
Copyright © 2023 Flavio Silva flavio1110@gmaill.com
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// ciCmd represents the ci command.
var ciCmd = &cobra.Command{
	Use:   "ci",
	Short: "Stage all modifications and commit them with the provided message",
	Long: `Stage all modifications (INCLUDING files not yet tracked) and commit them with the provided message. For example:

got ci "fix stuff"`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("missing commit messageZ. e.g. got nb feature/123")
		}

		path := getWorkDir()
		r, err := git.PlainOpen(path)
		exitIfError(err)

		w, err := r.Worktree()
		exitIfError(err)

		_, err = w.Add(".")
		exitIfError(err)

		_, err = w.Commit(args[0], &git.CommitOptions{})
		exitIfError(err)

		fmt.Println("Changes committed.")
	},
}

func init() {
	rootCmd.AddCommand(ciCmd)

}
