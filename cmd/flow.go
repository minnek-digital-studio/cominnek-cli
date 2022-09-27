package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/cmd/flow"
	"github.com/spf13/cobra"
)

var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "Manage git flow",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	flow.SetFlags()
	flowCmd.AddCommand(flow.FlowFeatureCmd)
	flowCmd.AddCommand(flow.FlowReleaseCmd)
	flowCmd.AddCommand(flow.FlowHotfixCmd)
	flowCmd.AddCommand(flow.FlowSupportCmd)
	rootCmd.AddCommand(flowCmd)
}