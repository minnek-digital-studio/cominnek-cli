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

func branchQuestion() {
	fmt.Println()

	if config.AppData.Branch.Type == "" {
		ask.One(&survey.Select{
			Message: "Select the branch type:",
			Options: []string{"feature", "bugfix", "hotfix", "release", "support"},
		}, &config.AppData.Branch.Type, survey.WithValidator(survey.Required))
	}

	if config.AppData.Branch.Ticket == "" {
		ask.One(&survey.Input{
			Message: "Enter the ticket number or name:",
		}, &config.AppData.Branch.Ticket, survey.WithValidator(survey.Required))
	}

	if !config.AppData.Branch.Stash && git_controller.CheckChanges() {
		ask.One(&survey.Confirm{
			Message: "Do you want to stash your changes?",
		}, &config.AppData.Branch.Stash)
	}
}

func Branch() {
	branchQuestion()
	ticket := config.AppData.Branch.Ticket

	if !config.AppData.Branch.Stash && git_controller.CheckChanges() {
		fmt.Println("You have uncommited changes, please commit them before creating a new branch")
		return
	}

	middleware(func() {
		switch config.AppData.Branch.Type {
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

		if config.AppData.Branch.Stash {
			git.Switch(originBranch)
			git.StashApply()
		}
	})

	if config.AppData.Branch.Stash {
		git.Stash("")
	}

	callBack()

	if config.AppData.Branch.Stash {
		git.StashApply()
	}
}
