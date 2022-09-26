package git

import (
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	checkers "github.com/Minnek-Digital-Studio/cominnek/pkg/git/Checkers"
)

func Feature(ticket string) {
	cmd := git_controller.Feature(ticket)
	checkers.CheckFlow(cmd)
}

func Fix(ticket string) {
	cmd := git_controller.Hotfix(ticket)
	checkers.CheckFlow(cmd)
}

func Release(ticket string) {
	cmd := git_controller.Release(ticket)
	checkers.CheckFlow(cmd)
}

func Support(ticket string) {
	cmd := git_controller.Support(ticket)
	checkers.CheckFlow(cmd)
}
