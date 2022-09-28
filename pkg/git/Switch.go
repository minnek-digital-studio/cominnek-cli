package git

import git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"

func Switch(branch string) {
	git_controller.Switch(branch)
}