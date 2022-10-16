package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/controllers"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Get the latest version of cominnek",
	Run: func(cmd *cobra.Command, args []string) {
		IgnoreCheckVersion = true
		controllers.Update()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
