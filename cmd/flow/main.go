package flow

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/spf13/cobra"
)

var stash bool

func SetFlags() {
	addFlags(FlowFeatureCmd)
	addFlags(FlowReleaseCmd)
	addFlags(FlowHotfixCmd)
	addFlags(FlowSupportCmd)
	addFlags(FlowBugfixCmd)
}

func addFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&stash, "stash", "s", false, "Stash changes before starting")
}

func setTicket(args []string) {
	if len(args) > 0 {
		config.AppData.Flow.Ticket = args[0]
	}
}
