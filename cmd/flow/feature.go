package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var FlowFeatureCmd = &cobra.Command{
	Use:   "feature <name>",
	Short: "create a new feature branch",
	Run: func(cmd *cobra.Command, args []string) {
		config.AppData.Flow.Type = "feature"
		config.AppData.Flow.Stash = stash
		setTicket(args)

		pkg_action.Flow()
	},
}
