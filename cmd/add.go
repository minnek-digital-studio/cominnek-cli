package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/extras"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [package]",
	Short: "Install a new package",
	Run: func(cmd *cobra.Command, args []string) {
		extras.AddPackage(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
