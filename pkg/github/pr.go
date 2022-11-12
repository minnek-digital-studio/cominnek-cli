package github

import (
	"log"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/controllers/app"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	github_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/github"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/fatih/color"
)

func _checkTicket(ticket string) string {
	if ticket == "" {
		loading.Stop()
		color.HiRed("This is branch is not linked to a ticket. Please use the -t flag to specify a ticket number")
		log.Fatal("Ticket number is required")
	}

	return ticket
}

func _getBranch(_baseBranch string, currentBranch string) string {
	baseBranch := _baseBranch

	if _baseBranch == "" {
		baseBranch = "develop"
	}

	return baseBranch
}

func __createPullRequest(_ticket, currentBranch, _baseBranch string) {
	origin := git_controller.GetOrigin()

	loading.Start("Preparing your pull request ")
	ticket := _checkTicket(_ticket)
	baseBranch := _getBranch(_baseBranch, currentBranch)
	body, title := git_controller.Pull_request(ticket, currentBranch, baseBranch)
	loading.Stop()

	github_controller.CreatePullRequest(github_controller.NewPullRequest{
		Title: title,
		Head:  currentBranch,
		Base:  baseBranch,
		Body:  body,
		Owner: origin.Owner,
		Repo:  origin.Repo,
		// Draft: true,
	})
}

func CreatePullRequest(_ticket string, _baseBranch string) {
	asserts := app.ConfigGlobal.PR.Asserts
	branch := git_controller.GetCurrentBranch()
	baseBranch := []string{_baseBranch}

	for _, assert := range asserts {
		if assert.Type == "*" {
			baseBranch = assert.BaseBranch
		}

		if strings.Contains(branch, assert.Type) {
			baseBranch = assert.BaseBranch
			break
		}
	}

	for _, _branch := range baseBranch {
		__createPullRequest(_ticket, branch, _branch)
	}

}
