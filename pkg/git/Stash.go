package git

import (
	"fmt"
	"log"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func _stash() {
	cmd := git_controller.Stash()
	loading.Start("Stashing changes ")
	err, out, errout := shell.Out(cmd)
	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		log.Fatal(errout)
	}

	loading.Stop()
	fmt.Println(out)
}

func StashApply() {
	cmd := git_controller.StashApply()
	loading.Start("Applying stashed changes ")
	err, out, errout := shell.Out(cmd)
	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		log.Fatal(errout)
	}

	loading.Stop()
	fmt.Println(out)
}

func Stash(branch string) {
	_stash()

	if branch != "" {
		loading.Start("Switching to " + branch + " ")

		Switch(branch)

		loading.Stop()
		StashApply()
	}
}
