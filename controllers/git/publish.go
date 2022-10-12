package git_controller

import "fmt"

func Publish(branch string) string {
	return fmt.Sprintf("git push --set-upstream origin %s", branch)
}
