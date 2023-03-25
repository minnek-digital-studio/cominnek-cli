package git_controller

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func Reset() {
	_, _, err := shell.Out("git reset -q HEAD -- .")

	if err != nil {
		panic(err)
	}
}


func ResetBy(r_type string, r_number string, r_commit string) string {
	cmd := "git reset"
	cmd += " --" + r_type

	if r_number != "" {
		cmd += " HEAD~" + r_number
	}

	if r_commit != "" {
		cmd += " " + r_commit
	}

	return cmd
}