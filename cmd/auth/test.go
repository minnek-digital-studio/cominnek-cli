package auth

import (
	"github.com/spf13/cobra"

	github_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/github"
)

var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "test connection to github",
	Run: func(cmd *cobra.Command, args []string) {
		github_controller.TestConnection()
	},
}
