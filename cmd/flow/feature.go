package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)
var FlowFeatureCmd = &cobra.Command{
	Use:   "feature <name>",
	Args: cobra.ExactArgs(1),
	Short: "create a new feature branch",
	Run: func(cmd *cobra.Command, args []string) {
		checker(args)
		
		exec := func() {
			git.Feature(args[0])
		}

		middleware(exec)
	},
}
