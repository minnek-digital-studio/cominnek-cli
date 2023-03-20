package pkg_action

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/folders"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
)

var currentDir bool

func createRepoQuestions() {
	ask.One(&survey.Confirm{
		Message: "Create repo in current directory?",
	}, &currentDir, nil)

	if currentDir {
	} else {
		ask.One(&survey.Input{
			Message: "Enter a name for the repo:",
			Help:   "This will create a new directory with the name you enter and create a new repo in that directory.",
		}, &config.AppData.CreateRepo.Name, survey.WithValidator(survey.Required))
	}
}

func CreateRepo() {
	createRepoQuestions()

	if config.AppData.CreateRepo.Name != "" {
		if !folders.CheckExists(config.AppData.CreateRepo.Name) {
			folders.Create(config.AppData.CreateRepo.Name)
		}
	}

	git.Create(config.AppData.CreateRepo.Name);
}