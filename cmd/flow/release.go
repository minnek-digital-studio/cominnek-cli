package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var FlowReleaseCmd = &cobra.Command{
	Use:   "release <name>",
	Args: cobra.ExactArgs(1),
	Short: "create a new release branch",
	Run: func(cmd *cobra.Command, args []string) {
		checker(args)

		exec := func() {
			git.Release(args[0])
		}
		middleware(exec)
	},
}
