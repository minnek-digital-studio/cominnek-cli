package cmd

import (
	"fmt"
	"os"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var stashCmd = &cobra.Command{
	Use:   "stash",
	Short: "Stash changes and switch to a branch",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			branch := args[0]
			git.Stash(branch)
		} else {
			fmt.Println("No branch provided")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(stashCmd)
}