/*
Copyright © 2023 Flavio Silva flavio1110@gmaill.com
*/
package cmd

import (
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push the current head to the remote",
	Long: `Push the current head to the remote origin. For example:

got push`,
	Run: func(cmd *cobra.Command, args []string) {
		path := getWorkDir()

		repo, err := git.PlainOpen(path)
		exitIfError(err)

		err = repo.Push(&git.PushOptions{
			RemoteName: "origin",
		})
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
