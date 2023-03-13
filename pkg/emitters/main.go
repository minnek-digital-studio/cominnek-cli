package emitters

import (
	"os"
	"path/filepath"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	github_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/github"
	emitterTypes "github.com/Minnek-Digital-Studio/cominnek/pkg/emitters/types"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
)

func RootEmitter() {
	pwd, _ := os.Getwd()
	events.App.Emit("init:root", &emitterTypes.IRootEmitter{
		Ticket: git_controller.GetTicketNumber(),
		Branch: git_controller.GetCurrentBranch(),
		User: emitterTypes.User{
			Name: *github_controller.GetCurrentUser().Name,
			Username: *github_controller.GetCurrentUser().Login,
		},
		Route:filepath.Dir(pwd) ,
	})
}
