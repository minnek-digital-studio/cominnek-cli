package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/cmd/auth"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/folders"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with Github",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	if !folders.CheckExists(config.Public.AppPath) {
		folders.Create(config.Public.AppPath)
	}

	authCmd.AddCommand(auth.LoginCmd)
	authCmd.AddCommand(auth.TestCmd)
	authCmd.AddCommand(auth.LogoutCmd)
	rootCmd.AddCommand(authCmd)
}
