package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/cmd/flow"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var flowCmd = &cobra.Command{
	Use:   "flow",
	Short: "Add a new repository to your Github account",
	Run: func(cmd *cobra.Command, args []string) {
		git.Add()
	},
}

func init() {
	flowCmd.AddCommand(flow.FlowFeatureCmd)
	rootCmd.AddCommand(flowCmd)
}