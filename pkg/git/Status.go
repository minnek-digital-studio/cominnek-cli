package git

import (
	"fmt"
	"log"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func Status() {
	loading.Start("Checking status of files ")

	out, errout, err := shell.Out(git_controller.Status())
	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		log.Fatal(errout)
	}

	loading.Stop()

	if out == "" {
		fmt.Println("No changes to commit")
	} else {
		fmt.Println(out)
	}
}

func ExistOnOrigin(branch string) bool {
	loading.Start("Checking if branch exists on origin ")
	out, _, err := shell.Out("git ls-remote --exit-code --heads origin " + branch)
	loading.Stop()

	if err != nil {
		return false
	}

	if out == "" {
		return false
	} else {
		return true
	}
}
