package git_controller

import (
	"fmt"
	"strings"
)

func getBaseCmd(branch string) string {
	return fmt.Sprintf("git branch %s; git checkout %s", branch, branch)
}

func Custom(path string, ticket string) (string, string) {
	branchName := strings.Replace(path, "*", ticket, 1)
	return getBaseCmd(branchName), branchName
}
