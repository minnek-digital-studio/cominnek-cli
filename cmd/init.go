package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/extras"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init a new package",
	Run: func(cmd *cobra.Command, args []string) {
		extras.InitProject()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
