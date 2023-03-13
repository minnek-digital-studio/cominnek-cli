package git_controller

import (
	"fmt"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func _removeTwoSpaces(str string) string {
	return strings.Replace(str, "  ", " ", -1)
}

func _getCommitMessage(msg string, ctype string, ticket string, scope string) string {
	commit_message := fmt.Sprintf("-m \"%v(%v): %v %v\"", ctype, scope, ticket, msg)

	if scope == "" {
		commit_message = fmt.Sprintf("-m \"%v: %v %v\"", ctype, ticket, msg)
	}

	return _removeTwoSpaces(commit_message)
}

func CheckChanges() bool {
	out, _, err := shell.Out(Status())

	if err != nil {
		return false
	}

	if out == "" {
		return false
	}

	return true
}

func Commit(msg string, _body string, ctype string, tiket string, scope string) (string, string) {
	trimScope := strings.TrimSpace(scope)
	body := strings.TrimSpace(_body)
	commit_message := _getCommitMessage(msg, ctype, tiket, trimScope)

	command := fmt.Sprintf("git commit %v", commit_message)

	if body != "" {
		command = fmt.Sprintf("git commit %v -m \"%v\"", commit_message, body)
	}

	return command, strings.Replace(commit_message, "-m ", "", 1)
}
