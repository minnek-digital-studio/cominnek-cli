package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var FlowReleaseCmd = &cobra.Command{
	Use:   "release <name>",
	Short: "create a new release branch",
	Run: func(cmd *cobra.Command, args []string) {
		config.AppData.Flow.Type = "release"
		config.AppData.Flow.Stash = stash
		setTicket(args)

		pkg_action.Flow()
	},
}
