package git

import (
	"os"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

var resetEmitter = new(emitters.Reset)

func Reset(r_type string, r_number string, r_commit string) {
	color.Yellow("\nResetting\n")
	cmd := git_controller.ResetBy(r_type, r_number, r_commit)
	_, _, err := shell.OutLive(cmd)

	if err != nil {
		resetEmitter.Failed(err.Error())
		os.Exit(1)
	}
}

func ResetAll() {
	cmd := git_controller.RemoveAllChanges()
	_, _, err := shell.Out(cmd)

	if err != nil {
		resetEmitter.Failed(err.Error())
		os.Exit(1)
	}
}
