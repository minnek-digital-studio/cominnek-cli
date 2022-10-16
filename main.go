package main

import (
	"github.com/Minnek-Digital-Studio/cominnek/cmd"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers"
)

func main() {
	config.Defaults()
	cmd.Execute()

	if !cmd.IgnoreCheckVersion {
		controllers.CheckUpdates(true)
	}
}
