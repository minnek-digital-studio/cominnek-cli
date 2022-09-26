package cmd

import (
	"fmt"

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
	},
}

func init() {
	publishCmd.PersistentFlags().StringVarP(&body, "body", "m", "", "Commit body message")
	publishCmd.PersistentFlags().StringVarP(&feat, "feature", "F", "{false}", "Add a new feature")
	publishCmd.PersistentFlags().StringVarP(&fix, "fix", "f", "{false}", "Fix an existing issue")
	publishCmd.PersistentFlags().StringVarP(&docs, "docs", "d", "{false}", "Add documentation")
	publishCmd.PersistentFlags().StringVarP(&refactor, "refactor", "r", "{false}", "Refactor an existing issue")
	publishCmd.PersistentFlags().StringVarP(&test, "test", "t", "{false}", "Add tests")
	publishCmd.PersistentFlags().StringVarP(&build, "build", "b", "{false}", "Build the project")
	rootCmd.AddCommand(publishCmd)
}
