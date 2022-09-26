package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var FlowReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "create a new release branch",
	Run: func(cmd *cobra.Command, args []string) {
		git.Release(args[0])
	},
}
