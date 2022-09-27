package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var FlowSupportCmd = &cobra.Command{
	Use:   "support <name>",
	Args: cobra.ExactArgs(1),
	Short: "create a new support branch",
	Run: func(cmd *cobra.Command, args []string) {
		checker(args)

		exec := func() {
			git.Support(args[0])
		}

		middleware(exec)
	},
}
