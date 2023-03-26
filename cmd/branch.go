package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/cmd/branch"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Create a new branch. This is variant of git-flow by Minnek",
	Run: func(cmd *cobra.Command, args []string) {
		pkg_action.Branch()
	},
}

func init() {
	branch.SetFlags()
	branchCmd.AddCommand(branch.BranchFeatureCmd)
	branchCmd.AddCommand(branch.BranchReleaseCmd)
	branchCmd.AddCommand(branch.BranchHotfixCmd)
	branchCmd.AddCommand(branch.BranchSupportCmd)
	branchCmd.AddCommand(branch.BranchBugfixCmd)
	branchCmd.AddCommand(branch.BranchTestCmd)
	rootCmd.AddCommand(branchCmd)
}
