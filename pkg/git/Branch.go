package git

import (
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	checkers "github.com/Minnek-Digital-Studio/cominnek/pkg/git/Checkers"
)

func Feature(ticket string) {
	cmd := git_controller.Feature(ticket)
	checkers.CheckBranch(cmd)
}

func Bugfix(ticket string) {
	cmd := git_controller.Bugfix(ticket)
	checkers.CheckBranch(cmd)
}

func HotFix(ticket string) {
	cmd := git_controller.Hotfix(ticket)
	checkers.CheckBranch(cmd)
}

func Release(ticket string) {
	cmd := git_controller.Release(ticket)
	checkers.CheckBranch(cmd)
}

func Support(ticket string) {
	cmd := git_controller.Support(ticket)
	checkers.CheckBranch(cmd)
}
