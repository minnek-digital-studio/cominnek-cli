package git_controller

import (
	"os"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/controllers/project"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

var SkipError bool = false

func GetCurrentBranch() string {
	cmd := "git rev-parse --abbrev-ref HEAD"
	out, _, err := shell.Out(cmd)

	if err != nil {
		if SkipError {
			return ""
		}

		color.Red("Something went wrong getting the current branch.")
		println("This may be because you don't have any commits yet, try making your first commit.")
		os.Exit(1)
	}

	return strings.TrimSpace(out)
}

func GetTicketNumber() string {
	currentBranch := GetCurrentBranch()

	var paths []string

	for _, path := range project.Config.Git.Branches {
		paths = append(paths, strings.ReplaceAll(path.Path, "*", ""))
	}

	for _, path := range paths {
		if strings.HasPrefix(currentBranch, path) {
			return strings.TrimPrefix(currentBranch, path)
		}
	}

	return ""
}

func GetBranchType() string {
	currentBranch := GetCurrentBranch()

	for _, path := range project.Config.Git.Branches {
		if strings.HasPrefix(currentBranch, strings.ReplaceAll(path.Path, "*", "")) {
			return path.Name
		}
	}

	return ""
}
