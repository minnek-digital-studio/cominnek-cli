package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/extras"
	"github.com/spf13/cobra"
)

var updateVersion = &cobra.Command{
	Use:   "update-version <version>",
	Short: "Create a commit for update version following conventional commits",
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		extras.UpdateVersion(version)
	},
}

func init() {
	rootCmd.AddCommand(updateVersion)
}
