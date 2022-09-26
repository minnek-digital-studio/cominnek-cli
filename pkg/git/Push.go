package git

import (
	"fmt"
	"log"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func _push() {
	cmd := git_controller.Push()
	err, out, errout := shell.Out(cmd)
	
	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		log.Fatal(errout)
	}

	fmt.Println(out)
}

func Push(msg string, body string, ctype string, scope string) {
	if msg != "" {
		Add()
		Status()
		Commit(msg, body, ctype, scope)
	}
	_push()
}
