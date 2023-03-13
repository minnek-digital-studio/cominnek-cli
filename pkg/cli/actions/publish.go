package pkg_action

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/github"
)

var publishEmmiter = new(emitters.Publish)

func publishQuestions() {
	if config.AppData.Publish.Ticket == "" {
		loading.Start("Reading branches...")
		ticket := git_controller.GetTicketNumber()
		loading.Stop()

		if ticket != "" {
			config.AppData.Publish.Ticket = ticket
			return
		}

		ask.One(&survey.Input{
			Message: "Enter the ticket number:",
		}, &config.AppData.Publish.Ticket, survey.WithValidator(survey.Required))
	}
}

func Publish() {
	publishEmmiter.Init()
	if !config.AppData.Publish.IgnoreCommit {
		config.AppData.Publish.IgnoreCommit = !git_controller.CheckChanges()
	}

	if !config.AppData.Publish.IgnoreCommit {
		Commit(false)
	}

	pushQuestion()
	publishQuestions()

	if !config.AppData.Publish.IgnoreCommit {
		executeCommit()
	}

	github.Publish(config.AppData.Publish.Ticket)

	if config.AppData.Push.Merge != "" {
		git.Merge(config.AppData.Push.Merge)
	}
}
