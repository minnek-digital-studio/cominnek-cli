package git_controller

import "github.com/Minnek-Digital-Studio/cominnek/pkg/shell"

func Reset() {
	_, _, err := shell.Out("git reset -q HEAD -- .")

	if err != nil {
		panic(err)
	}
}
