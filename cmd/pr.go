package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/github"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Create a new pull request",
	Run: func(cmd *cobra.Command, args []string) {
		github.CreatePullRequest()
	},
}

func init() {
	rootCmd.AddCommand(prCmd)
}