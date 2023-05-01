package pkg_action

import (
	"fmt"
	"os"

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

var branchEmitter = new(emitters.Branch)

func branchQuestion() {
	fmt.Println()

	if config.AppData.Branch.Type == "" {
		ask.One(&survey.Select{
			Message: "Select the branch type:",
			Options: []string{"feature", "bugfix", "hotfix", "release", "support", "test", "sync"},
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

	branchEmitter.Init(data)

	if !config.AppData.Branch.Stash && git_controller.CheckChanges() {
		fmt.Println("You have uncommitted changes, please commit them before creating a new branch")
		branchEmitter.Failed(emitterTypes.IBranchFailedData{
			Error: "Uncommitted changes",
			Data:  data,
		})
		return
	}

	middleware(func(exec bool) string {
		var branch string;
		switch config.AppData.Branch.Type {
		case "feature":
			branch = git.Feature(ticket, exec)
		case "bugfix":
			branch = git.Bugfix(ticket, exec)
		case "hotfix":
			branch = git.HotFix(ticket, exec)
		case "release":
			branch = git.Release(ticket, exec)
		case "support":
			branch = git.Support(ticket, exec)
		case "test":
			branch = git.Test(ticket, exec)
		case "sync":
			branch = git.Sync(ticket, exec)
		}

		return branch
	})
}

func middleware(callBack func(exe bool) string) {
	loading.Start("Getting current branch...")
	originBranch := git_controller.GetCurrentBranch()
	loading.Stop()

	branch := callBack(false)

	if git_controller.CheckBranchExist(branch) {
		color.Red("Branch already exists")
		branchEmitter.Failed(emitterTypes.IBranchFailedData{
			Error: "Branch already exists",
			Data: emitterTypes.IBranchEventData{
				Ticket: config.AppData.Branch.Ticket,
				Type:   config.AppData.Branch.Type,
			},
		})
		os.Exit(1)
	}

	events.App.On("cleanup", func(payload ...interface{}) {
		originErr := payload[0].(string)
		fmt.Println("Cleaning up")

		if config.AppData.Branch.Stash {
			git.Switch(originBranch)
			git.StashApply()
		}

		color.Red("Branch creation failed")
		branchEmitter.Failed(emitterTypes.IBranchFailedData{
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

	callBack(true)

	color.Green("Branch created successfully")

	if config.AppData.Branch.Stash {
		git.StashApply()
	}
}
