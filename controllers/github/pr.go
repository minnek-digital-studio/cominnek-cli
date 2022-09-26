package github_controller

import (
	"fmt"

	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/google/go-github/v47/github"
)

type NewPullRequest struct {
	Title               string
	Head                string
	Base                string
	Body                string
	Owner				string
	Repo				string
	Draft               bool
}

func CreatePullRequest(prData NewPullRequest) {
	loading.Start("Creating pull request ")
	client := client()
	user, _, err := client.Users.Get(ctx, "")

	if err != nil {
		fmt.Println(err)
		return
	}

	newPR := &github.NewPullRequest{
		Title:               github.String(prData.Title),
		Head:                github.String(prData.Head),
		Base:                github.String(prData.Base),
		Body:                github.String(prData.Body),
		MaintainerCanModify: github.Bool(true),
		Draft:               github.Bool(prData.Draft),
	}

	pr, _, err := client.PullRequests.Create(ctx, prData.Owner, prData.Repo, newPR)
	if err != nil {
		fmt.Println(err)
		return
	}

	client.Issues.AddAssignees(ctx, prData.Owner, prData.Repo, *pr.Number, []string{user.GetLogin()})

	if err != nil {
		fmt.Println(err)
		return
	}

	loading.Stop()

	fmt.Printf("PR created: %s\n", pr.GetHTMLURL())
}
