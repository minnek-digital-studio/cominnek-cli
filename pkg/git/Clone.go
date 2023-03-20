package git

import (
	"os"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func Clone(url string)  { 
	cmd := git_controller.Clone(url)
	_, _, err := shell.OutLive(cmd)

	if err != nil {
		println(color.HiRedString("Someting went wrong"))
		os.Exit(1)
	}

	color.Green("Cloned repository successfully")
}
