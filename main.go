package main

import (
	"github.com/Minnek-Digital-Studio/cominnek/cmd"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/logger"
	"github.com/Minnek-Digital-Studio/cominnek/helper"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	"github.com/fatih/color"
)

func init() {
	events.Watcher()
}

func main() {
	helper.PrintName()
	logger.PrintLn(color.HiRedString("!!!!!!!!!!!!!!!!!!!!!!!!!!!!"))
	logger.PrintLn(color.HiRedString("!!!!!"), color.HiYellowString("Logs Are Enable"), color.HiRedString("!!!!!!"))
	logger.PrintLn(color.HiRedString("!!!!!!!!!!!!!!!!!!!!!!!!!!!!"))
	config.Defaults()
	
	c := make(chan bool)

	go func() {
		emitters.RootEmitter()
		c <- true
	}()

	cmd.Execute()

	if !cmd.IgnoreCheckVersion {
		controllers.CheckUpdates(true)
	}

	<-c
}
