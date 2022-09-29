package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/github"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Create a new pull request",
	Run: func(cmd *cobra.Command, args []string) {
		github.CreatePullRequest(ticket)
	},
}

func init() {
	prCmd.PersistentFlags().StringVarP(&ticket, "ticket", "t", "", "Ticket number")
	rootCmd.AddCommand(prCmd)
}
