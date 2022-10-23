package git

import (
	"fmt"
	"log"
	"strings"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func _getCurrentBranch() string {
	loading.Start("Merging")
	branch := git_controller.GetCurrentBranch()
	loading.Stop()
	return branch
}

func _mergeErr(currentBranch string, branch string, out string) {
	loading.Start("Error merging " + currentBranch + " into " + branch)
	conflictFiles := git_controller.Conflict{}.GetFiles()
	loading.Stop()

	if strings.Contains(out, "CONFLICT") {
		fmt.Println("There are conflicts in the following files:")
		color.HiRed(conflictFiles)
		fmt.Print("\n\n")
		log.Fatal("Conflict detected. Please resolve the conflict and try again.")
	}

	fmt.Println("There was an error merging " + currentBranch + " into " + branch)
	fmt.Println(out)
	log.Fatal("Error merging branch")
}

func _merge(currentBranch string, branch string) string {
	loading.Start("Merging " + currentBranch + " into " + branch)
	cmd := git_controller.Merge(currentBranch)
	out, _, err := shell.Out(cmd)

	if err != nil {
		loading.Stop()
		_mergeErr(currentBranch, branch, out)
	}

	loading.Stop()

	return out
}

func Merge(branch string) {
	currentBranch := _getCurrentBranch()
	git_controller.Switch(branch)

	pkg.AppEvent.On("cleanup", func(...interface{}) {
		fmt.Println("Cleaning up")
		fmt.Println("You have some conflicts to resolve. After you have resolved them, run the following command to continue:")
		fmt.Println("git merge " + currentBranch)
	})

	git_controller.Pull()

	out := _merge(currentBranch, branch)

	fmt.Println(out)
	log.Println("Merged " + currentBranch + " into " + branch)
}
