package cli

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
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
		ask.One(&survey.Select{
			Message: "What do you want to do?",
			Options: actions,
		}, &config.AppData.Action, nil)
	}

	if config.AppData.Action == "Exit" {
		os.Exit(0)
	}

	if config.AppData.Action == "Commit" {
		pkg_action.Commit(true)
	}

	if config.AppData.Action == "Push" {
		pkg_action.Push()
	}

	if config.AppData.Action == "Publish" {
		pkg_action.Publish()
	}

	if config.AppData.Action == "Pull Request" {
		pkg_action.PullRequest()
	}
}
