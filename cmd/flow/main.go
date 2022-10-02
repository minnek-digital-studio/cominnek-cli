package flow

import (
	"fmt"
	"os"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
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

func checker(args []string) {
	if len(args) < 1 || args[0] == "" {
		fmt.Println("No branch name provided")
		os.Exit(1)
	}
}

func middleware(callBack func()) {
	loading.Start("Checking for uncommitted changes ")
	originBranch := git_controller.GetCurrentBranch()
	loading.Stop()

	pkg.App.On("cleanup", func(...interface{}) {
		fmt.Println("Cleaning up")

		if stash {
			git.Switch(originBranch)
			git.StashApply()
		}
	})

	if stash {
		git.Stash("")
	}

	callBack()

	if stash {
		git.StashApply()
	}
}
