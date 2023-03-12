package pkg_action

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/github"
)

func prQuestions() {
	if config.AppData.PullRequest.Ticket == "" {
		loading.Start("Reading branches...")
		ticket := git_controller.GetTicketNumber()
		loading.Stop()

		if ticket != "" {
			config.AppData.PullRequest.Ticket = ticket
			return
		}

		ask.One(&survey.Input{
			Message: "Enter the ticket number:",
		}, &config.AppData.PullRequest.Ticket, survey.WithValidator(survey.Required))
	}

	if config.AppData.PullRequest.Base == "" {
		ask.One(&survey.Select{
			Message: "Select the base branch:",
			Options: []string{"auto", "develop", "master"},
			Default: "auto",
		}, &config.AppData.PullRequest.Base, survey.WithValidator(survey.Required))

		if config.AppData.PullRequest.Base == "auto" {
			config.AppData.PullRequest.Base = ""
		}
	}
}

func PullRequest() {
	prQuestions()

	github.NewCreatePullRequest(config.AppData.PullRequest.Ticket, config.AppData.PullRequest.Base)
}
