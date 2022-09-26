package git_controller

import (
	"fmt"
	"log"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func GetCurrentBranch() string {
	cmd := "git rev-parse --abbrev-ref HEAD"
	err, out, errout := shell.Out(cmd)

	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		log.Fatal(errout)
	}

	return strings.TrimSpace(out)
}

func GetTicketNumber() string {
	currentBranch := GetCurrentBranch();

	if strings.HasPrefix(currentBranch, "feature/") {
		return strings.TrimPrefix(currentBranch, "feature/")
	} else if strings.HasPrefix(currentBranch, "hotfix/") {
		return strings.TrimPrefix(currentBranch, "hotfix/")
	} else if strings.HasPrefix(currentBranch, "release/") {
		return strings.TrimPrefix(currentBranch, "release/")
	} else if strings.HasPrefix(currentBranch, "support/") {
		return strings.TrimPrefix(currentBranch, "support/")
	}

	return ""
}