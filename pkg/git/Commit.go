package git

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/controllers"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

var commiteEmmiter = new(emitters.Commit)

func _commit(msg string, body string, ctype string, scope string, ticket string) string {
	color.Yellow("\nCommiting files\n")
	cmd, message := git_controller.Commit(msg, body, ctype, ticket, scope)
	out, _, err := shell.OutLive(cmd)

	if err != nil {
		loading.Stop()

		if strings.Contains(out, "nothing to commit") {
			fmt.Println("\nAborting commit...")
			commiteEmmiter.Failed("Nothing to commit")
			os.Exit(1)
		} else {
			commiteEmmiter.Failed(out)
			log.Fatal("Commit failed")
		}
	}

	commiteEmmiter.Success(message)

	return out
}

func _checkTicket(ticket string) string {
	if ticket == "" {
		loading.Stop()
		if !controllers.Confirm("No ticket number found. Commit anyway?", false) {
			fmt.Println("Aborting commit")
			commiteEmmiter.Failed("Aborted by user")
			os.Exit(0)
		}

		loading.Start("Commiting files ")
	}

	return ticket
}

func Commit(msg string, body string, ctype string, scope string) {
	loading.Start("Commiting files ")
	currentBranch := git_controller.GetCurrentBranch()

	if strings.HasPrefix(currentBranch, "bugfix/") {
		if ctype == "feat" {
			errorMsg := "Bugfix branch cannot have a feature commit"
			loading.Stop()
			color.HiRed("Error:")
			log.Fatal(errorMsg)

			commiteEmmiter.Failed(errorMsg)

			os.Exit(1)
		}
	}

	ticket := _checkTicket(git_controller.GetTicketNumber())
	loading.Stop()
	_commit(msg, body, ctype, scope, ticket)

}

func CommitWithoutTicket(msg string, body string, ctype string, scope string) {
	_commit(msg, body, ctype, scope, "")
}
