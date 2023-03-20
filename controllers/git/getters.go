package git_controller

import (
	"os"
	"strings"

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

	if strings.HasPrefix(currentBranch, "feature/") {
		return strings.TrimPrefix(currentBranch, "feature/")
	} else if strings.HasPrefix(currentBranch, "hotfix/") {
		return strings.TrimPrefix(currentBranch, "hotfix/")
	} else if strings.HasPrefix(currentBranch, "release/") {
		return strings.TrimPrefix(currentBranch, "release/")
	} else if strings.HasPrefix(currentBranch, "support/") {
		return strings.TrimPrefix(currentBranch, "support/")
	} else if strings.HasPrefix(currentBranch, "bugfix/") {
		return strings.TrimPrefix(currentBranch, "bugfix/")
	}

	return ""
}
