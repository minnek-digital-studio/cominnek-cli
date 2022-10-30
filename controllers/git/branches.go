package git_controller

import (
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func ListBranches() []string {
	branches, _, err := shell.Out(`git branch --format "%(refname:short)"`)

	if err != nil {
		panic(err)
	}

	branches = strings.ReplaceAll(branches, " ", "")
	branches = strings.ReplaceAll(branches, "\n", ",")
	branches = strings.Trim(branches, ",")
	return strings.Split(branches, ",")
}
