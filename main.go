package main

import (
	"github.com/Minnek-Digital-Studio/cominnek/cmd"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/app"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/logger"
	"github.com/Minnek-Digital-Studio/cominnek/helper"
	"github.com/fatih/color"
)

func main() {
	app.ConfigReader()
	helper.PrintName()

	logger.PrintLn(color.HiRedString("!!!!!!!!!!!!!!!!!!!!!!!!!!!!"))
	logger.PrintLn(color.HiRedString("!!!!!"), color.HiYellowString("Logs Are Enable"), color.HiRedString("!!!!!!"))
	logger.PrintLn(color.HiRedString("!!!!!!!!!!!!!!!!!!!!!!!!!!!!"))

	config.Defaults()
	cmd.Execute()

	if !cmd.IgnoreCheckVersion {
		controllers.CheckUpdates(true)
	}
}
