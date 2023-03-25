package cmd

import (
	"os"
	"strconv"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var hardReset bool
var softReset bool
var mixedReset bool
var mergeReset bool
var keepReset bool

var resetCmd = &cobra.Command{
	Use:   "reset <number>",
	Short: "Reset changes at the current branch.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			cmd.Help()
			return
		}

		getArgs(args)
		getResetType()
		pkg_action.Reset()
	},
}

func init() {
	resetCmd.Flags().BoolVar(&hardReset, "hard", false, "Reset HEAD, index and working tree")
	resetCmd.Flags().BoolVar(&softReset, "soft", false, "Reset only HEAD")
	resetCmd.Flags().BoolVar(&mixedReset, "mixed", false, "Reset HEAD and index")
	resetCmd.Flags().BoolVar(&mergeReset, "merge", false, "Reset HEAD, index and working tree")
	resetCmd.Flags().BoolVar(&keepReset, "keep", false, "Reset HEAD, index and working tree")
	resetCmd.Flags().StringVarP(&config.AppData.Reset.Number, "number", "n", "", "The number of commits to reset.")
	resetCmd.Flags().StringVarP(&config.AppData.Reset.Commit, "commit", "c", "", "The commit hash to reset to.")
	resetCmd.Flags().StringVarP(&config.AppData.Reset.Target, "target", "r", "", "The target to reset to.")
	resetCmd.Flags().BoolVarP(&config.AppData.Reset.Confirm, "confirm", "y", false, "Confirm the reset.")
	rootCmd.AddCommand(resetCmd)
}

func getResetType() {
	if config.AppData.Reset.Type == "" {
		switch {
		case hardReset:
			config.AppData.Reset.Type = "hard"
		case softReset:
			config.AppData.Reset.Type = "soft"
		case mixedReset:
			config.AppData.Reset.Type = "mixed"
		case mergeReset:
			config.AppData.Reset.Type = "merge"
		case keepReset:
			config.AppData.Reset.Type = "keep"
		}
	}
}

func getArgs(args []string) {
	if len(args) == 1 {
		arg := args[0]
		_, err := strconv.Atoi(arg)
		if err != nil {
			byCommit(arg)
		} else {
			config.AppData.Reset.Number = arg
		}
	}
}

func byCommit(arg string) {
	if git.ValidateCommitHash(arg) {
		commitInfo := git.GetCommitByHash(arg)
		config.AppData.Reset.Commit = git.GetCommitHash(commitInfo)
		color.Yellow("⚠ Resetting to commit: ")
		println(commitInfo)
	} else {
		color.Red("Invalid commit hash or number of commits.")
		os.Exit(1)
	}
}
