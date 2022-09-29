package github

import (
	"fmt"
	"log"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	github_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/github"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/fatih/color"
)

func _checkTicket(ticket string) string {
	_ticket := git_controller.GetTicketNumber()
	if ticket != "" && _ticket != "" {
		loading.Stop()
		fmt.Println("This branch is linked to a ticket, you can't use the -t flag")
		log.Fatal("ticket number: ", _ticket)
	}

	if _ticket != "" {
		ticket = _ticket
	}

	if ticket == "" {
		loading.Stop()
		color.HiRed("This is branch is not linked to a ticket. Please use the -t flag to specify a ticket number")
		log.Fatal("Ticket number is required")
	}

	return ticket
}

func CreatePullRequest(_ticket string) {
	origin := git_controller.GetOrigin()

	loading.Start("Preparing your pull request ")
	currentBranch := git_controller.GetCurrentBranch()
	ticket := _checkTicket(_ticket)
	body := git_controller.Pull_request(ticket)
	title := currentBranch
	loading.Stop()

	github_controller.CreatePullRequest(github_controller.NewPullRequest{
		Title: title,
		Head:  currentBranch,
		Base:  "develop",
		Body:  body,
		Owner: origin.Owner,
		Repo:  origin.Repo,
		Draft: true,
	})
}
