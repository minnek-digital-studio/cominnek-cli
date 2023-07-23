package branch

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/project"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
	"github.com/spf13/cobra"
)

var stash bool

func SetCommands(Command *cobra.Command) {
	project.ReadConfigFile()

	for _, branch := range project.Config.Git.Branches {
		if branch.Config.Hidden {
			continue
		}

		branchName := branch.Name
		branchData := branch
		branchCmd := &cobra.Command{
			Use:     branchName,
			Example: branchName + " <name>",
			Short:   "Create a new " + branchName + " branch",
			Long:    branch.Config.Description,
			Run: func(cmd *cobra.Command, args []string) {
				setTicket(args)
				config.AppData.Branch.Stash = stash
				config.AppData.Branch.Data = branchData
				pkg_action.Branch()
			},
		}

		addFlags(branchCmd)
		Command.AddCommand(branchCmd)
	}
}

func addFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolVarP(&stash, "stash", "s", false, "Stash changes before starting")
}

func setTicket(args []string) {
	if len(args) > 0 {
		config.AppData.Branch.Ticket = args[0]
	}
}
