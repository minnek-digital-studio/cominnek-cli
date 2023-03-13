package github_controller

import (
	"fmt"
	"os"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	"github.com/fatih/color"
	"github.com/google/go-github/v47/github"
)

var pullRequestEmmiter = new(emitters.PullRequest)

type NewPullRequest struct {
	Title string
	Head  string
	Base  string
	Body  string
	Owner string
	Repo  string
	Draft bool
}

func showExistingPR(prData NewPullRequest) {
	client := client()

	existing_pr, _, err := client.PullRequests.List(ctx, prData.Owner, prData.Repo, &github.PullRequestListOptions{
		State: "open",
		Head:  prData.Head,
		Base: prData.Base,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nExisting PR:")

	for _, pr := range existing_pr {
		fmt.Println("\t" + pr.GetHTMLURL())
	}
	fmt.Println()
}

func CreatePullRequest(prData NewPullRequest) {
	loading.Start("Creating pull request ")
	client := client()
	user := GetCurrentUser()

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
		loading.Stop()
		color.Red("Something went wrong:")

		if strings.Contains(err.Error(), "Code:custom Message:") {
			customMessage := strings.Split(err.Error(), "Code:custom Message:")[1]
			clearMessageS1 := strings.Replace(customMessage, "}", "", -1)
			clearMessageS2 := strings.Replace(clearMessageS1, "]", "", -1)
			message := fmt.Sprintf("\t%s", clearMessageS2)
			fmt.Println(message)

			if strings.Contains(message, "A pull request already exists for") {
				showExistingPR(prData)
				pullRequestEmmiter.Failed("A pull request already exists for")
			} else {
				pullRequestEmmiter.Failed(message)
			}

			os.Exit(1)
		}

		if strings.Contains(err.Error(), "Field:head") {
			message := fmt.Sprintf("%v branch does not exist on remote", prData.Head)
			fmt.Println("\t" + message)
			fmt.Println("\n" + "use: 'cominnek publish'")
			pullRequestEmmiter.Failed(message)
			os.Exit(1)
		}

		fmt.Println(err)
		pullRequestEmmiter.Failed(err.Error())
		os.Exit(1)
	}

	client.Issues.AddAssignees(ctx, prData.Owner, prData.Repo, *pr.Number, []string{user.GetLogin()})

	if err != nil {
		fmt.Println(err)
		pullRequestEmmiter.Failed(err.Error())
		return
	}

	loading.Stop()
	pullRequestEmmiter.Success(pr.GetHTMLURL())
	color.Yellow("\nPull request created\n")
	println(pr.GetHTMLURL(), "\n")
}
