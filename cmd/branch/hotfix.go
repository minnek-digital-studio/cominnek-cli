package branch

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var BranchHotfixCmd = &cobra.Command{
	Use:   "hotfix <name>",
	Short: "create a new hotfix branch from master",
	Run: func(cmd *cobra.Command, args []string) {
		config.AppData.Branch.Type = "hotfix"
		config.AppData.Branch.Stash = stash
		setTicket(args)

		pkg_action.Branch()
	},
}
