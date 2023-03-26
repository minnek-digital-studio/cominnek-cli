package branch

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var BranchReleaseCmd = &cobra.Command{
	Use:   "release <name>",
	Short: "create a new release branch from develop",
	Run: func(cmd *cobra.Command, args []string) {
		config.AppData.Branch.Type = "release"
		config.AppData.Branch.Stash = stash
		setTicket(args)

		pkg_action.Branch()
	},
}
