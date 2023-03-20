package cli

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	pkg_action "github.com/Minnek-Digital-Studio/cominnek/pkg/cli/actions"
)

func askActionsRepo() {
	actions := []string{
		"Clone Repo",
		"Create Repo",
		"Exit",
	}

	if config.AppData.Action == "" {
		ask.One(&survey.Select{
			Message: "What do you want to do?",
			Options: actions,
		}, &config.AppData.Action, nil)
	}

	if config.AppData.Action == "Exit" {
		os.Exit(0)
	}

	if config.AppData.Action == "Clone Repo" {
		pkg_action.Clone()
	}
}
