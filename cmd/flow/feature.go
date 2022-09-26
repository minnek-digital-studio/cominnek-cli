package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var FlowFeatureCmd = &cobra.Command{
	Use:   "feature",
	Short: "create a new feature branch",
	Run: func(cmd *cobra.Command, args []string) {
		git.Feature(args[0])
	},
}
