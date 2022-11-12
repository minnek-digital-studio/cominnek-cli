package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var skipCommit bool

var pushCmd = &cobra.Command{
	Use:   "push <message>",
	Short: "push a branch to GitHub",
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
		config.AppData.Push.IgnoreCommit = skipCommit

		pkg_action.Push()
	},
}

func init() {
	pushCmd.Flags().BoolVar(&skipCommit, "skip-commit", false, "Skip the commit and only push the branch")
	AddFlags{}.Push(pushCmd)
	rootCmd.AddCommand(pushCmd)
}
