package auth

import (
	github_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/github"
	"github.com/spf13/cobra"
)

var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout from your github account",
	Run: func(cmd *cobra.Command, args []string) {
		github_controller.Logout()
	},
}
