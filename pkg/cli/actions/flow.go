package pkg_action

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
)

func flowQuestion() {
	if config.AppData.Flow.Type == "" {
		ask.One(&survey.Select{
			Message: "Select the branch type:",
			Options: []string{"feature", "bugfix", "hotfix", "release", "support"},
		}, &config.AppData.Flow.Type, survey.WithValidator(survey.Required))
	}

	if config.AppData.Flow.Ticket == "" {
		ask.One(&survey.Input{
			Message: "Enter the ticket number or name:",
		}, &config.AppData.Flow.Ticket, survey.WithValidator(survey.Required))
	}
}

func Flow() {
	flowQuestion()
	ticket := config.AppData.Flow.Ticket

	middleware(func() {
		switch config.AppData.Flow.Type {
		case "feature":
			git.Feature(ticket)
		case "bugfix":
			git.Bugfix(ticket)
		case "hotfix":
			git.HotFix(ticket)
		case "release":
			git.Release(ticket)
		case "support":
			git.Support(ticket)
		}
	})
}

func middleware(callBack func()) {
	loading.Start("Getting current branch...")
	originBranch := git_controller.GetCurrentBranch()
	loading.Stop()

	events.App.On("cleanup", func(...interface{}) {
		fmt.Println("Cleaning up")

		if config.AppData.Flow.Stash {
			git.Switch(originBranch)
			git.StashApply()
		}
	})

	if config.AppData.Flow.Stash {
		git.Stash("")
	}

	callBack()

	if config.AppData.Flow.Stash {
		git.StashApply()
	}
}
