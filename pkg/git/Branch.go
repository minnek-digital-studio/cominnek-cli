package git

import (
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	checkers "github.com/Minnek-Digital-Studio/cominnek/pkg/git/Checkers"
)

func Feature(ticket string, execute bool) string {
	cmd, branch := git_controller.Feature(ticket)
	if(execute) {
		checkers.CheckBranch(cmd)
	}
	return branch
}

func Bugfix(ticket string, execute bool) string {
	cmd, branch := git_controller.Bugfix(ticket)
	if(execute) {
		checkers.CheckBranch(cmd)
	}
	return branch
}

func HotFix(ticket string, execute bool) string {
	cmd, branch := git_controller.Hotfix(ticket)
	if(execute) {
		checkers.CheckBranch(cmd)
	}
	return branch
}

func Release(ticket string, execute bool) string {
	cmd, branch := git_controller.Release(ticket)
	if(execute) {
		checkers.CheckBranch(cmd)
	}
	return branch
}

func Support(ticket string, execute bool) string {
	cmd, branch := git_controller.Support(ticket)
	if(execute) {
		checkers.CheckBranch(cmd)
	}
	return branch
}

func Test(ticket string, execute bool) string {
	cmd, branch := git_controller.Test(ticket)
	if(execute) {
		checkers.CheckBranch(cmd)
	}
	return branch
}
