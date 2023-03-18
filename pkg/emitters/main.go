package emitters

import (
	"os"
	"path/filepath"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	emitterTypes "github.com/Minnek-Digital-Studio/cominnek/pkg/emitters/types"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
)

func RootEmitter() {
	if !git_controller.CheckGitRepo() {
		return
	}

	pwd, _ := os.Getwd()
	events.App.Emit("init:root", &emitterTypes.IRootEmitter{
		Ticket: git_controller.GetTicketNumber(),
		Branch: git_controller.GetCurrentBranch(),
		//TODO: find a way to execute this only if is necessary
		// User: emitterTypes.User{
		// 	Name: *github_controller.GetCurrentUser().Name,
		// 	Username: *github_controller.GetCurrentUser().Login,
		// },
		Route: filepath.Dir(pwd),
	})
}
