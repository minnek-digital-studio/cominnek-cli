package git

import (
	"log"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

var pushEmmiter = new(emitters.Push)

func _push() {
	color.Yellow("\nPushing to remote\n")
	cmd := git_controller.Push()
	out, outErr, err := shell.OutLive(cmd)

	if err != nil {
		if outErr != "" {
			pushEmmiter.Failed(outErr)
		}

		if out != "" {
			pushEmmiter.Failed(out)
		}

		log.Fatal("Error pushing to remote")
	}

	pushEmmiter.Success("Pushed to remote")
}

func Push() {
	_push()
	log.Println("Push complete")
}

func PushWithOutTicket(msg string, body string, ctype string, scope string) {
	if msg != "" {
		Add()
		Status()
		CommitWithoutTicket(msg, body, ctype, scope)
	}
	_push()
	log.Println("Push complete")
}
