package git

import (
	"log"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func _push() {
	color.Yellow("\nPushing to remote\n")
	cmd := git_controller.Push()
	_, _, err := shell.OutLive(cmd)

	if err != nil {
		log.Fatal("Error pushing to remote")
	}

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
