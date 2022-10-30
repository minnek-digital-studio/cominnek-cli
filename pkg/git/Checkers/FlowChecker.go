package checkers

import (
	"fmt"
	"log"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func FetchData() {
	loading.Start("Checking Origin ")
	fetch := shell.ExecuteCommand("git fetch origin", false)

	if fetch != "" {
		loading.Update("Data updated")
		loading.Stop()
		fmt.Println(fetch)
	}

}

func GetChanges() {
	if git_controller.CheckChangesFromOrigin() {
		color.YellowString("\n\nThere are changes from origin.\n")
		git_controller.Pull()
	}
}

func CheckFlow(mainCmd string) {
	branch := "develop"

	if !git_controller.CheckIfBranch(branch) {
		git_controller.Switch(branch)
	}

	color.HiYellowString("Fetching data from origin...")
	FetchData()

	GetChanges()

	loading.Start("Starting new flow ")

	out, errout, err := shell.Out(mainCmd)
	if err != nil {
		loading.Stop()
		fmt.Println(out)
		fmt.Println(errout)

		events.App.Emit("cleanup")

		log.Fatal(errout)
	}

	loading.Stop()

	fmt.Println(out)
}
