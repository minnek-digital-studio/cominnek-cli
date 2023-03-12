package branch

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var BranchTestCmd = &cobra.Command{
	Use:   "test <name>",
	Short: "create a new test branch from develop",
	Run: func(cmd *cobra.Command, args []string) {
		config.AppData.Branch.Type = "test"
		config.AppData.Branch.Stash = stash
		setTicket(args)

		pkg_action.Branch()
	},
}