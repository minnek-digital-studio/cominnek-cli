package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge <branch>",
	Args: cobra.ExactArgs(1),
	Short: "Merge a branch into the current branch",
	Long: `Merge a branch into the current branch. This command will
merge the current branch into the branch specified. This command
will not work if there are any conflicts. If there are conflicts,
you will need to resolve them before running this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		branch := args[0]
		git.Merge(branch)
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}