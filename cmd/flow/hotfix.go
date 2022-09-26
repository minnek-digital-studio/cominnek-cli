package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var FlowHotfixCmd = &cobra.Command{
	Use:   "hotfix",
	Short: "create a new hotfix branch",
	Run: func(cmd *cobra.Command, args []string) {
		git.Fix(args[0])
	},
}
