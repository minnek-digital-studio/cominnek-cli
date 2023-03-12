package git

import (
	"log"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func PushPublish() {
	color.Yellow("\nPushing to remote\n")
	currentBranch := git_controller.GetCurrentBranch()
	cmd := git_controller.Publish(currentBranch)
	_, _, err := shell.OutLive(cmd)

	if err != nil {
		log.Fatal("Error pushing to remote")
	}
}
