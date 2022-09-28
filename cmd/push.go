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

		if merge != "" {
			git.Merge(merge)
		}
	},
}

func init() {
	AddFlags{}.Push(pushCmd)
	rootCmd.AddCommand(pushCmd)
}
