package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var FlowBugfixCmd = &cobra.Command{
	Use:   "bugfix <name>",
	Args:  cobra.ExactArgs(1),
	Short: "create a new feature branch",
	Run: func(cmd *cobra.Command, args []string) {
		checker(args)

		exec := func() {
			git.Bugfix(args[0])
		}

		middleware(exec)
	},
}
