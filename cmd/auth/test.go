package auth

import (
	"github.com/spf13/cobra"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
)

var TestCmd = &cobra.Command{
	Use:   "test",
	Short: "test connection to github",
	Run: func(cmd *cobra.Command, args []string) {
		// github_controller.TestConnection()	
		git_controller.GetOrigin()
	},
}
