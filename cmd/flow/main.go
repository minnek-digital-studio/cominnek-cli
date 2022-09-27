package flow

import (
	"fmt"
	"os"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var stash bool

func SetFlags() {
	addFlags(FlowFeatureCmd)
	addFlags(FlowReleaseCmd)
	addFlags(FlowHotfixCmd)
	addFlags(FlowSupportCmd)
}

func addFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&stash, "stash", "s", false, "Stash changes before starting")
}

func checker(args []string) {
	if len(args) < 1 || args[0] == "" {
		fmt.Println("No branch name provided")
		os.Exit(1)
	}
}

func middleware(callBack func()) {
	if stash {
		git.Stash("")
	}

	callBack()

	if stash {
		git.StashApply()
	}
}
