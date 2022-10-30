package git_controller

import (
	"fmt"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

// create a class
type Conflict struct {
}

func (x Conflict) GetFiles() string {
	cmd := "git diff --name-only --diff-filter=U --relative"

	out, _, err := shell.Out(cmd)

	if err != nil {
		fmt.Println("Error getting conflicts files")
	}

	return out
}
