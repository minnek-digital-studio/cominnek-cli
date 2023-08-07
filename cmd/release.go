package cmd

import (
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "release a new version",
	Run: func(cmd *cobra.Command, args []string) {
		pkg_action.Release()
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}
