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

	err, out, errout := shell.Out(git_controller.Status())
	if err != nil {
		fmt.Println(out)
		fmt.Println(errout)
		log.Fatal(errout)
	}

	loading.Stop()
	
	if(out == "") {
		fmt.Println("No changes to commit")
	} else {	
		fmt.Println(out)
	}
}
