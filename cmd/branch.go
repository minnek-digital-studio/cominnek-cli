package cmd

import (
	"os"

	"github.com/Minnek-Digital-Studio/cominnek/cmd/branch"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/project"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/cli"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Create a new branch.",
	Annotations: map[string]string{
		"command": "branch",
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !cli.CheckConfig() {
			color.Red("\nSorry, you need to initialize the project first.")
			os.Exit(1)
		}

		project.ReadConfigFile(true)
		pkg_action.Branch()
	},
}

func init() {
	branch.SetCommands(branchCmd)
	rootCmd.AddCommand(branchCmd)
}
