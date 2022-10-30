package pkg_action

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
)

func stashQuestions() {
	if config.AppData.Stash.Branch == "" {
		loading.Start("Reading branches...")
		branches := git_controller.ListBranches()
		currentBranch := git_controller.GetCurrentBranch()

		for i, branch := range branches {
			if branch == currentBranch {
				branches = append(branches[:i], branches[i+1:]...)
				break
			}
		}

		loading.Stop()

		ask.One(&survey.Select{
			Message: "Select a branch to merge your changes into:",
			Options: branches,
		}, &config.AppData.Stash.Branch, survey.WithValidator(survey.Required))
	}
}

func Stash() {
	stashQuestions()
	git.Stash(config.AppData.Stash.Branch)
}
