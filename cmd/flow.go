package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/cmd/flow"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "Manage git flow",
	Run: func(cmd *cobra.Command, args []string) {
		pkg_action.Flow()
	},
}

func init() {
	flow.SetFlags()
	flowCmd.AddCommand(flow.FlowFeatureCmd)
	flowCmd.AddCommand(flow.FlowReleaseCmd)
	flowCmd.AddCommand(flow.FlowHotfixCmd)
	flowCmd.AddCommand(flow.FlowSupportCmd)
	flowCmd.AddCommand(flow.FlowBugfixCmd)
	rootCmd.AddCommand(flowCmd)
}
