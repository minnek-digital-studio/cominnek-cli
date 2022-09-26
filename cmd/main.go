package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cominnek",
	Short: "Manage your git flow using the conventional commit standard",
	Version: "2.0.0",
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