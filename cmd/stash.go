package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var stashCmd = &cobra.Command{
	Use:   "stash <branch>",
	Short: "Stash changes and switch to a branch",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			config.AppData.Stash.Branch = args[0]
		}

		pkg_action.Stash()
	},
}

func init() {
	rootCmd.AddCommand(stashCmd)
}
