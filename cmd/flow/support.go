package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var FlowSupportCmd = &cobra.Command{
	Use:   "support",
	Short: "create a new support branch",
	Run: func(cmd *cobra.Command, args []string) {
		git.Support(args[0])
	},
}
