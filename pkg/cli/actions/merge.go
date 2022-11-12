package pkg_action

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
)

func mergeQuestions() {
	if config.AppData.Merge.Branch == "" {
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
		}, &config.AppData.Merge.Branch, survey.WithValidator(survey.Required))
	}
}

func Merge() {
	mergeQuestions()
	git.Merge(config.AppData.Merge.Branch)

}
