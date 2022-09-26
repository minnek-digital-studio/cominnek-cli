package checkers

import (
	"fmt"
	"time"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func fetchData(s *spinner.Spinner) {
	s.Start()
	s.Prefix = "Checking Origin..."
	fetch := shell.ExecuteCommand("git fetch origin", false)

	if fetch != "" {
		s.Prefix = "Data updated"
		fmt.Println(fetch)
	}
	s.Stop()
}

func CheckFlow(mainCmd string) {
	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	branch := "develop"

	if !git_controller.CheckIfBranch(branch) {
		color.HiBlue("\tSwitching to develop branch...")
		shell.ExecuteCommand("git checkout "+branch, false)
		color.HiGreen("\tSwitched to develop\n")
	}

	color.HiYellowString("Fetching data from origin...")

	fetchData(s)

	if git_controller.CheckChangesFromOrigin() {
		color.YellowString("\n\nThere are changes from origin.\n")
		s.UpdateCharSet(spinner.CharSets[2])
		s.Start()
		s.Prefix = "Pulling lastest changes from origin..."
		fmt.Print("\n\n")
		shell.ExecuteCommand(git_controller.Pull())
		s.Stop()
	}

	shell.ExecuteCommand(mainCmd)
}
