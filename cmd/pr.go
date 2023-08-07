package cmd

import (
	"os"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/project"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/cli"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Create a new pull request",
	Run: func(cmd *cobra.Command, args []string) {
		if !cli.CheckConfig() {
			color.Red("\nSorry, you need to initialize the project first.")
			os.Exit(1)
		}
		project.ReadConfigFile(true)

		config.AppData.PullRequest.Ticket = ticket
		config.AppData.PullRequest.Base = baseBranch

		pkg_action.PullRequest()
	},
}

func init() {
	prCmd.PersistentFlags().StringVarP(&ticket, "ticket", "t", "", "Ticket number")
	prCmd.PersistentFlags().StringVarP(&baseBranch, "base", "b", "develop", "Base branch")
	rootCmd.AddCommand(prCmd)
}
