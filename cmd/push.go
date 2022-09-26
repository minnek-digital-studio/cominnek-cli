package cmd

import (
	"fmt"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push <message>",
	Short: "push a branch to GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		msg := ""
		scope := ""

		if len(args) > 0 {
			msg = args[0]
		}

		if msg == "" {
			scope = getScope(true)
			fmt.Println("Not commit message provided")
			fmt.Print("Commit aborted\n\n")
		}

		if msg != "" {
			scope = getScope(false)
		}

		git.Push(msg, body, ctype, scope)
	},
}

func init() {
	pushCmd.PersistentFlags().StringVarP(&body, "body", "m", "", "Commit body message")
	pushCmd.PersistentFlags().StringVarP(&feat, "feature", "F", "{false}", "Add a new feature")
	pushCmd.PersistentFlags().StringVarP(&fix, "fix", "f", "{false}", "Fix an existing issue")
	pushCmd.PersistentFlags().StringVarP(&docs, "docs", "d", "{false}", "Add documentation")
	pushCmd.PersistentFlags().StringVarP(&refactor, "refactor", "r", "{false}", "Refactor an existing issue")
	pushCmd.PersistentFlags().StringVarP(&test, "test", "t", "{false}", "Add tests")
	pushCmd.PersistentFlags().StringVarP(&build, "build", "b", "{false}", "Build the project")
	rootCmd.AddCommand(pushCmd)
}
