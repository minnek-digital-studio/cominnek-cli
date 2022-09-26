package git

import (
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	checkers "github.com/Minnek-Digital-Studio/cominnek/pkg/git/Checkers"
)

func Feature(ticket string) {
	cmd := git_controller.Feature(ticket)
	checkers.CheckFlow(cmd)
}
