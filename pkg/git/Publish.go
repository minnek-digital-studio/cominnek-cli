package git

import (
	"fmt"
	"log"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func PushPublish() {
	loading.Start("Pushing to remote ")
	currentBranch := git_controller.GetCurrentBranch()
	cmd := git_controller.Publish(currentBranch)

	err, out, errout := shell.Out(cmd)
	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		log.Fatal(errout)
	}

	loading.Stop()
	fmt.Println(out)
}
