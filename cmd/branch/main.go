package branch

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/spf13/cobra"
)

var stash bool

func SetFlags() {
	addFlags(BranchFeatureCmd)
	addFlags(BranchReleaseCmd)
	addFlags(BranchHotfixCmd)
	addFlags(BranchSupportCmd)
	addFlags(BranchBugfixCmd)
}

func addFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&stash, "stash", "s", false, "Stash changes before starting")
}

func setTicket(args []string) {
	if len(args) > 0 {
		config.AppData.Branch.Ticket = args[0]
	}
}
