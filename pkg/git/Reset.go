package git

import (
	"os"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func Reset(r_type string, r_number string, r_commit string) {
	color.Yellow("\nResetting\n")
	cmd := git_controller.ResetBy(r_type, r_number, r_commit)
	_, _, err := shell.OutLive(cmd)
	
	if err != nil {
		os.Exit(1)
	}
}