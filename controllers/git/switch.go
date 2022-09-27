package git_controller

import (
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func Switch(branch string) {
	switchingMsg := "Switching to " + branch + " branch "
	switchedMsg := "\tSwitched to " + branch + " branch\n"
	loading.Start(switchingMsg)

	shell.ExecuteCommand("git checkout "+branch, false)
	loading.Stop()
	color.HiGreen(switchedMsg)
}