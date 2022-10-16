package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/cmd/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configure your Github account",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	configCmd.AddCommand(config.PRConfigCmd)
	rootCmd.AddCommand(configCmd)
}