package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var FlowSupportCmd = &cobra.Command{
	Use:   "support <name>",
	Short: "create a new support branch",
	Run: func(cmd *cobra.Command, args []string) {
		config.AppData.Flow.Type = "support"
		config.AppData.Flow.Stash = stash
		setTicket(args)

		pkg_action.Flow()
	},
}
