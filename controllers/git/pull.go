package git_controller

import (
	"fmt"
	"log"

	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func Pull() {
	cmd := "git pull"

	loading.Start("Pulling changes from origin ")
	fmt.Print("\n\n")

	out, errout, err := shell.Out(cmd)
	if err != nil {
		loading.Stop()
		fmt.Println(out)
		fmt.Println(errout)

		pkg.AppEvent.Emit("cleanup")

		log.Fatal(errout)
	}

	loading.Stop()
	fmt.Println(out)
}
