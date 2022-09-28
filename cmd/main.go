package cmd

import (
	"fmt"
	"os"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cominnek",
	Short: "Manage your git flow using the conventional commit standard",
	Version: config.Public.Version,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cominnek is a CLI for Github")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rootCmd.SetVersionTemplate("cominnek version {{.Version}}\n")
}