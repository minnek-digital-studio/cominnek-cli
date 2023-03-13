package pkg_action

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	emitterTypes "github.com/Minnek-Digital-Studio/cominnek/pkg/emitters/types"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/fatih/color"
)

var branchEmmiter = new(emitters.Branch)

func branchQuestion() {
	fmt.Println()

	if config.AppData.Branch.Type == "" {
		ask.One(&survey.Select{
			Message: "Select the branch type:",
			Options: []string{"feature", "bugfix", "hotfix", "release", "support", "test"},
		}, &config.AppData.Branch.Type, survey.WithValidator(survey.Required))
	}

	if config.AppData.Branch.Ticket == "" {
		message := "Enter the ticket number or name:"

		if config.AppData.Branch.Type == "release" {
			message = "Enter the release version:"
		}

		if config.AppData.Branch.Type == "hotfix" {
			message = "Enter the hotfix version:"
		}

		if config.AppData.Branch.Type == "test" {
			message = "Enter the test version:"
		}

		ask.One(&survey.Input{
			Message: message,
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
	data := emitterTypes.IBranchEventData{
		Type:   config.AppData.Branch.Type,
		Ticket: ticket,
	}

	branchEmmiter.Init(data)

	if !config.AppData.Branch.Stash && git_controller.CheckChanges() {
		fmt.Println("You have uncommited changes, please commit them before creating a new branch")
		branchEmmiter.Failed(emitterTypes.IBranchFailedData{
			Error: "Uncommited changes",
			Data:  data,
		})
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
		case "test":
			git.Test(ticket)
		}
	})
}

func middleware(callBack func()) {
	loading.Start("Getting current branch...")
	originBranch := git_controller.GetCurrentBranch()
	loading.Stop()

	events.App.On("cleanup", func(payload ...interface{}) {
		originErr := payload[0].(string)
		fmt.Println("Cleaning up")

		if config.AppData.Branch.Stash {
			git.Switch(originBranch)
			git.StashApply()
		}

		color.Red("Branch creation failed")
		branchEmmiter.Failed(emitterTypes.IBranchFailedData{
			Error: "Branch creation failed: " + originErr,
			Data: emitterTypes.IBranchEventData{
				Ticket: config.AppData.Branch.Ticket,
				Type:   config.AppData.Branch.Type,
			},
		})
	})

	if config.AppData.Branch.Stash {
		git.Stash("")
	}

	callBack()

	color.Green("Branch created successfully")

	if config.AppData.Branch.Stash {
		git.StashApply()
	}
}
