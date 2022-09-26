package auth

import (
	"github.com/spf13/cobra"

	github_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/github"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login into your github account",
	Run: func(cmd *cobra.Command, args []string) {
		github_controller.Login()	
	},
}
