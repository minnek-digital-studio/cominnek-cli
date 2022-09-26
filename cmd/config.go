package cmd

import (
	"log"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var email, name string

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configure your Github account",
	Run: func(cmd *cobra.Command, args []string) {

		if name == "" {
			log.Fatal("No name specified")
		}

		if email == "" {
			log.Fatal("No email specified")
		}

		git.Config(name, email)
	},
}

func init() {
	configCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "Your name")
	configCmd.PersistentFlags().StringVarP(&email, "email", "e", "", "Your email")
	rootCmd.AddCommand(configCmd)
}