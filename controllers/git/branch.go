package git_controller

import "fmt"

func getBaseCmd(branch string) string {
	return fmt.Sprintf("git branch %s; git checkout %s", branch, branch)
}

func Feature(ticket string) string {
	return getBaseCmd("feature/" + ticket)
}

func Bugfix(ticket string) string {
	return getBaseCmd("bugfix/" + ticket)
}

func Hotfix(ticket string) string {
	return getBaseCmd("hotfix/" + ticket)
}

func Release(ticket string) string {
	return getBaseCmd("release/" + ticket)
}

func Support(ticket string) string {
	return getBaseCmd("support/" + ticket)
}

func Test(ticket string) string {
	return getBaseCmd("test/" + ticket)
}
