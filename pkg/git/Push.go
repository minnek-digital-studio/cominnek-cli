package git

import (
	"fmt"
	"log"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func _push() {
	loading.Start("Pushing to remote ")
	cmd := git_controller.Push()
	out, errout, err := shell.Out(cmd)

	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		log.Fatal(errout)
	}

	loading.Stop()
	fmt.Println(out)
}

func Push(msg string, body string, ctype string, scope string) {
	if msg != "" {
		Add()
		Status()
		Commit(msg, body, ctype, scope)
	}
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
