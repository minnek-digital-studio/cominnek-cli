package main

import (
	"time"

	"github.com/Minnek-Digital-Studio/cominnek/cmd"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/logger"
	"github.com/Minnek-Digital-Studio/cominnek/helper"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	"github.com/fatih/color"
)

func init() {
	events.Watcher()
	config.AppData.Start = time.Now()

	github_template := ".github/PULL_REQUEST_TEMPLATE.md"
	if files.CheckExist(github_template) {
		config.Public.PRBody = github_template
	}
}

func main() {
	verChan := make(chan bool)
	emitChan := make(chan bool)

	helper.PrintName()

	loading.Start("Loading internal modules")

	if !controllers.Connected() {
		loading.Stop()
		color.Yellow("We are not able to connect to the internet, some features may not work")
	}

	loading.Stop()

	logger.PrintLn(color.HiRedString("!!!!!!!!!!!!!!!!!!!!!!!!!!!!"))
	logger.PrintLn(color.HiRedString("!!!!!"), color.HiYellowString("Logs Are Enable"), color.HiRedString("!!!!!!"))
	logger.PrintLn(color.HiRedString("!!!!!!!!!!!!!!!!!!!!!!!!!!!!"))
	config.Defaults()

	go func() {
		if !cmd.IgnoreCheckVersion {
			controllers.CheckUpdates()
		}
		verChan <- true
	}()

	go func() {
		emitters.RootEmitter()
		emitChan <- true
	}()

	cmd.Execute()

	if !cmd.IgnoreCheckVersion && <-verChan {
		controllers.PrintUpdateMessage()
	}

	<-emitChan

	println("Done in", time.Since(config.AppData.Start).String())
}
