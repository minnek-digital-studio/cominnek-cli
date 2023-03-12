package branch

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var BranchSupportCmd = &cobra.Command{
	Use:   "support <name>",
	Short: "create a new support branch from master",
	Run: func(cmd *cobra.Command, args []string) {
		config.AppData.Branch.Type = "support"
		config.AppData.Branch.Stash = stash
		setTicket(args)

		pkg_action.Branch()
	},
}
