package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new repository to your Github account",
	Run: func(cmd *cobra.Command, args []string) {
		git.Add()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}