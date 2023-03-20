package git

import (
	"os"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func Create(folder string) {
	var cmd string;

	if folder != "" {
		cmd = "cd " + folder;
	}

	cmd += ";" + git_controller.CreateRepo();

	_, _, err := shell.OutLive(cmd);


	if err != nil {
		color.Red("Something went wrong creating the repo.")
		os.Exit(1)
	}

	color.Green("Repo created successfully.")
}