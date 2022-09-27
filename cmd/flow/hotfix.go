package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var FlowHotfixCmd = &cobra.Command{
	Use:   "hotfix <name>",
	Args: cobra.ExactArgs(1),
	Short: "create a new hotfix branch",
	Run: func(cmd *cobra.Command, args []string) {
		checker(args)
		
		exec := func() {
			git.Fix(args[0])
		}
		middleware(exec)
	},
}
