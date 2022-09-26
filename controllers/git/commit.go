package git_controller

import (
	"fmt"
	"strings"
)

func Commit(msg string, _body string, ctype string, tiket string, scope string) string {
	trimScope := strings.TrimSpace(scope)
	body := strings.TrimSpace(_body)
	commit_message := fmt.Sprintf("-m \"%v(%v): %v %v\"", ctype, trimScope, tiket, msg)

	if trimScope == "" {
		commit_message = fmt.Sprintf("-m \"%v: %v %v\"", ctype, tiket, msg)
	}

	command := fmt.Sprintf("git commit %v", commit_message)

	if body != "" {
		command = fmt.Sprintf("git commit %v -m \"%v\"", commit_message, body)
	}

	return command
}
