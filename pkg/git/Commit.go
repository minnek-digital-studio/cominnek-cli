package git

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/controllers"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func Commit(msg string, body string, ctype string, scope string) {
	loading.Start("Commiting files ")
	ticket := git_controller.GetTicketNumber()

	if ticket == "" {
		if !controllers.Confirm("No ticket number found. Commit anyway?", "n") {
			loading.Stop()
			fmt.Println("Aborting commit")
			return
		}
	}
	cmd := git_controller.Commit(msg, body, ctype, ticket, scope)
	err, out, errout := shell.Out(cmd)

	if err != nil {
		loading.Stop()

		if strings.Contains(out, "nothing to commit") {
			fmt.Println(out)
			fmt.Println("Aborting commit...")

			os.Exit(1)
		} else {
			fmt.Println("Error: ", err)
			fmt.Println(out)
			fmt.Println(errout)
			log.Fatal(errout)
		}
	}

	loading.Stop()
	fmt.Println(out)
}
