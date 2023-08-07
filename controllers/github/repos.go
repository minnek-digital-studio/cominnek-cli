package github_controller

import (
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/google/go-github/v47/github"
)

func GetRepoList() (repos []*github.Repository) {
	client := client()
	loading.Start("Getting Repo List")

	repos, _, err := client.Repositories.List(ctx, "", &github.RepositoryListOptions{
		Sort: "updated",
		ListOptions: github.ListOptions{
			PerPage: 100,
		},
	})

	if err != nil {
		loading.Stop()
		println("Sorry cannot get repo list")
	}

	loading.Stop()

	return repos
}

func GetRepoURL(owner string, repo string) string {
	client := client()

	repository, _, err := client.Repositories.Get(ctx, owner, repo)

	if err != nil {
		println("Sorry cannot get repo list")
	}

	return repository.GetHTMLURL()
}
