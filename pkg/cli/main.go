package cli

import (
	"github.com/AlecAivazis/survey/v2"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/extras"
	"github.com/fatih/color"
)

func Main() {
	if git_controller.CheckGitRepo() {
		if CheckConfig() {
			askActions()
		}

		return
	}

	askActionsRepo()
}

func CheckConfig() bool {
	if extras.CheckIfConfigExists() == false {
		var initProject bool
		color.HiYellow("We have detected that you have not initialized the project yet.")

		ask.One(&survey.Confirm{
			Message: "Do you want to initialize the project?",
		}, &initProject, nil)

		if initProject {
			extras.InitProject()
		}

		return initProject
	}

	return true
}
