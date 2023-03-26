package cli

import git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"

func Main() {
	if git_controller.CheckGitRepo() {
		askActions()
		return
	}

	askActionsRepo()
}
