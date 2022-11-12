package cmd

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge <branch>",
	Short: "Merge a branch into the current branch",
	Long: `Merge a branch into the current branch. This command will
merge the current branch into the branch specified. This command
will not work if there are any conflicts. If there are conflicts,
you will need to resolve them before running this command.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			config.AppData.Merge.Branch = args[0]
		}

		pkg_action.Merge()
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}
