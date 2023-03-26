package git_controller

import "fmt"

func getBaseCmd(branch string) string {
	return fmt.Sprintf("git branch %s; git checkout %s", branch, branch)
}

func Feature(ticket string) (string, string) {
	branchName := "feature/" + ticket
	return getBaseCmd(branchName), branchName
}

func Bugfix(ticket string) (string, string) {
	branchName := "bugfix/" + ticket
	return getBaseCmd(branchName), branchName
}

func Hotfix(ticket string) (string, string) {
	branchName := "hotfix/" + ticket
	return getBaseCmd(branchName), branchName
}

func Release(ticket string) (string, string) {
	branchName := "release/" + ticket
	return getBaseCmd(branchName), branchName
}

func Support(ticket string) (string, string) {
	branchName := "support/" + ticket
	return getBaseCmd(branchName), branchName
}

func Test(ticket string) (string, string) {
	branchName := "test/" + ticket
	return getBaseCmd(branchName), branchName
}
