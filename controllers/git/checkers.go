package git_controller

import (
	"fmt"
	"log"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func CheckIfBranch(branch string) bool {
	currentBranch := GetCurrentBranch()

	if currentBranch != branch {
		fmt.Println("This is not " + branch + " branch")
		return false
	}

	return true
}

func CheckChangesFromOrigin() bool {
	out, errout, err := shell.Out("git status")

	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		log.Fatal(errout)
	}

	ready := strings.Contains(out, "Your branch is up to date")
	return !ready
}

func CheckBranchExist(branch string) bool {
	_, _, err := shell.Out("git rev-parse --verify " + branch)
	return err == nil
}
