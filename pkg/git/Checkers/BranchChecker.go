package checkers

import (
	"fmt"
	"log"
	"os"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	emitterTypes "github.com/Minnek-Digital-Studio/cominnek/pkg/emitters/types"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

var branchEmitter = new(emitters.Branch)

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

func CheckBranch(mainCmd string) {
	branchData := config.AppData.Branch.Data
	branch := branchData.From

	if !git_controller.CheckBranchExist(branch) {
		println("\nBranch " + branch + " not found\n")
		os.Exit(1)
	}

	if !git_controller.CheckIfBranch(branch) {
		git_controller.Switch(branch)
	}

	color.HiYellowString("Fetching data from origin...")
	FetchData()

	GetChanges()

	loading.Start("Creating a new branch ")

	out, errout, err := shell.Out(mainCmd)
	if err != nil {
		loading.Stop()
		fmt.Println(out)
		fmt.Println(errout)

		events.App.Emit("cleanup", err.Error())

		log.Fatal(errout)
	}

	loading.Stop()

	branchEmitter.Success(emitterTypes.IBranchEventData{
		Type:   config.AppData.Branch.Data.Name,
		Ticket: config.AppData.Branch.Ticket,
	})

	fmt.Println(out)
}
