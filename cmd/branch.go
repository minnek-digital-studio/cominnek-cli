package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/cmd/branch"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/project"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Create a new branch.",
	Run: func(cmd *cobra.Command, args []string) {
		project.ReadConfigFile(true)

		pkg_action.Branch()
	},
}

func init() {
	branch.SetCommands(branchCmd)
	rootCmd.AddCommand(branchCmd)
}
