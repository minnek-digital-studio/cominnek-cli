package cmd

import (
	"fmt"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/github"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish <message>",
	Short: "Publish a branch to GitHub and create a pull request as Draft",
	Run: func(cmd *cobra.Command, args []string) {
		scope := ""
		msg := message[0]
		body := ""

		if len(message) > 1 {
			body = message[1]
		}

		if msg == "" {
			scope = getScope(true)
			fmt.Println("Not commit message provided")
			fmt.Println("Commit aborted")
			fmt.Println("Starting publish...")
		}

		if msg != "" {
			scope = getScope(false)
		}

		github.Publish(msg, body, ctype, scope, ticket)

		if merge != "" {
			git.Merge(merge)
		}
	},
}

func init() {
	AddFlags{}.Push(publishCmd)
	publishCmd.PersistentFlags().StringVarP(&ticket, "ticket", "t", "", "Ticket number")
	rootCmd.AddCommand(publishCmd)
}
