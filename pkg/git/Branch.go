package git

import (
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	checkers "github.com/Minnek-Digital-Studio/cominnek/pkg/git/Checkers"
)

func Custom(path string, ticket string, execute bool) string {
	cmd, branch := git_controller.Custom(path, ticket)
	if execute {
		checkers.CheckBranch(cmd)
	}
	return branch
}
