package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish <message>",
	Short: "Publish a branch to GitHub and create a pull request as Draft",
	Run: func(cmd *cobra.Command, args []string) {
		msg := ""
		body := ""

		if len(message) > 0 {
			msg = message[0]
		}

		if len(message) > 1 {
			body = message[1]
		}

		config.AppData.Commit.AddAll = addAll
		config.AppData.Commit.Message = msg
		config.AppData.Commit.Scope = getScope(true)
		config.AppData.Commit.Type = ctype
		config.AppData.Commit.Body = body
		config.AppData.Push.Merge = merge
		config.AppData.Publish.Ticket = ticket
		config.AppData.Publish.IgnoreCommit = skipCommit

		pkg_action.Publish()

		if merge != "" {
			git.Merge(merge)
		}
	},
}

func init() {
	AddFlags{}.Push(publishCmd)
	publishCmd.PersistentFlags().StringVarP(&ticket, "ticket", "t", "", "Ticket number")
	publishCmd.Flags().BoolVar(&skipCommit, "skip-commit", false, "Skip the commit and only push the branch")
	rootCmd.AddCommand(publishCmd)
}
