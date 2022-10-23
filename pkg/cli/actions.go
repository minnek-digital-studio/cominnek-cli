package cli

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
)

func askActions() {
	actions := []string{
		"Commit",
		"Create Branch",
		"Pull Request",
		"Publish",
		"Push",
		"Exit",
	}

	if config.AppData.Action == "" {
		survey.AskOne(&survey.Select{
			Message: "What do you want to do?",
			Options: actions,
		}, &config.AppData.Action, nil)
	}

	if config.AppData.Action == "Exit" {
		os.Exit(0)
	}

	if config.AppData.Action == "Commit" {
		pkg_action.Commit()
	}
}
