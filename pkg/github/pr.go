package github

import (
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	github_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/github"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
)

func CreatePullRequest() {
	origin := git_controller.GetOrigin();

	loading.Start("Preparing your pull request ")
	currentBranch := git_controller.GetCurrentBranch();
	ticket := git_controller.GetTicketNumber();
	body := git_controller.Pull_request(ticket);
	title := currentBranch;
	loading.Stop()

	github_controller.CreatePullRequest(github_controller.NewPullRequest{
		Title: title,
		Head: currentBranch,
		Base: "develop",
		Body: body,
		Owner: origin.Owner,
		Repo: origin.Repo,
		Draft: true,
	})
}