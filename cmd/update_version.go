package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/extras"
	"github.com/spf13/cobra"
)

var updateVersion = &cobra.Command{
	Use:   "update-version <version>",
	Short: "BigCommerce update version",
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		extras.UpdateVersion(version)
	},
}

func init() {
	rootCmd.AddCommand(updateVersion)
}
