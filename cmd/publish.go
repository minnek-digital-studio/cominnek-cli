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
		msg := ""
		scope := ""

		if len(args) > 0 {
			msg = args[0]
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

		github.Publish(msg, body, ctype, scope)

		if merge != "" {
			git.Merge(merge)
		}
	},
}

func init() {
	AddFlags{}.Push(publishCmd)
	rootCmd.AddCommand(publishCmd)
}
