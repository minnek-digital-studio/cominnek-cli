package git

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/controllers"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

var commitEmitter = new(emitters.Commit)

func _commit(msg string, body string, ctype string, scope string, ticket string) string {
	color.Yellow("\nCommitting files\n")
	cmd, message := git_controller.Commit(msg, body, ctype, ticket, scope)
	out, _, err := shell.OutLive(cmd)

	if err != nil {
		loading.Stop()

		if strings.Contains(out, "nothing to commit") {
			fmt.Println("\nAborting commit...")
			commitEmitter.Failed("Nothing to commit")
			os.Exit(1)
		} else {
			commitEmitter.Failed(out)
			log.Fatal("Commit failed")
		}
	}

	commitEmitter.Success(message)

	return out
}

func _checkTicket(ticket string) string {
	if ticket == "" {
		loading.Stop()
		if !controllers.Confirm("No ticket number found. Commit anyway?", false) {
			fmt.Println("Aborting commit")
			commitEmitter.Failed("Aborted by user")
			os.Exit(0)
		}

		loading.Start("Committing files ")
	}

	return ticket
}

func Commit(msg string, body string, ctype string, scope string) {
	loading.Start("Committing files ")
	currentBranch := git_controller.GetCurrentBranch()

	if strings.HasPrefix(currentBranch, "bugfix/") {
		if ctype == "feat" {
			errorMsg := "Bugfix branch cannot have a feature commit"
			loading.Stop()
			color.HiRed("Error:")
			log.Fatal(errorMsg)

			commitEmitter.Failed(errorMsg)

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

func GetAllCommits() []string {
	out, _, err := shell.Out("git log --pretty=format:'%h: %s - %an, %ar'")

	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(out, "\n")
}

func GetAllCommitsForRelease(lastReleaseHash string) []string {
	cmd := "git log --pretty=format:'message[%s];hash[%h]'"

	if lastReleaseHash != "" {
		cmd += " " + lastReleaseHash + "..HEAD"
	}

	out, _, err := shell.Out(cmd)

	if err != nil {
		println("Error getting commits")
		os.Exit(1)
	}

	return strings.Split(out, "\n")
}

func LastTag() string {
	tags := GetTags()
	fmt.Println("Tags:", tags)

	for tag := range tags {
		if !strings.Contains(tags[tag], "-") {
			return tags[tag]
		}
	}

	return ""
}

func GetTags() []string {
	out, _, err := shell.Out("git tag --sort=-creatordate")

	if err != nil {
		println("Error getting tags")
		return []string{}
	}

	return strings.Split(out, "\n")
}

func GetCommitHash(msg string) string {
	return strings.Split(msg, ":")[0]
}

func ValidateCommitHash(hash string) bool {
	_, _, err := shell.Out("git rev-parse --verify " + hash)

	return err == nil
}

func GetCommitByHash(hash string) string {
	os := runtime.GOOS
	grepCmd := "grep"

	if os == "windows" {
		grepCmd = "Select-String"
	}

	cmd := fmt.Sprintf("git log --pretty=oneline --pretty=format:'%%h: %%s' %s | "+grepCmd+" %s", hash, hash)
	out, _, err := shell.Out(cmd)

	if err != nil {
		log.Fatal(err)
	}

	return out
}
