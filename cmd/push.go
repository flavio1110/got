/*
Copyright © 2023 Flavio Silva flavio1110@gmaill.com
*/
package cmd

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// pushCmd represents the push command.
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push the current head to the remote",
	Long: `Push the current head to the remote origin. For example:

got push 
or
got push f `,
	Run: func(cmd *cobra.Command, args []string) {
		path := getWorkDir()

		repo, err := git.PlainOpen(path)
		exitIfError(err)

		force := len(args) > 0 && args[0] == "f"

		err = repo.Push(&git.PushOptions{
			RemoteName: "origin",
			Force:      force,
		})
		if err != nil {
			if errors.Is(err, git.NoErrAlreadyUpToDate) {
				fmt.Println("Nothing to push. Remote is already up to date!")
				return
			}

			exitIfError(err)
		}
		ref, err := repo.Head()
		exitIfError(err)

		fmt.Printf("Changes pushed with success to %s\n", ref.Name())
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
