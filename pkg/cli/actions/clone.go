package pkg_action

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	github_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/github"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
)

func cloneQuestions() {
	if config.AppData.Clone.Url == "" {
		repos := github_controller.GetRepoList()
		repoMap := make(map[string]string)
		var options []string = []string{
			"Custom URL",
		}

		for _, repo := range repos {
			url := *repo.SSHURL
			name := repo.GetName()
			options = append(options, name)
			repoMap[name] = url
		}

		prompt := &survey.Select{
			Message:  "Which repo do you want to clone?",
			Options:  options,
			PageSize: 10,
			VimMode:  true,
			Help:     "Select the repo you want to clone",
		}

		ask.One(prompt, &config.AppData.Clone.Url)

		if config.AppData.Clone.Url == "Custom URL" {
			ask.One(&survey.Input{
				Message: "Enter the URL of the repo you want to clone",
			}, &config.AppData.Clone.Url)
		} else {
			config.AppData.Clone.Url = repoMap[config.AppData.Clone.Url]
		}
	}
}

func Clone() {
	cloneQuestions()
	git.Clone(config.AppData.Clone.Url)
}
