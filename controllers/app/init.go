package app

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/folders"
)

func Init() {
	if !folders.CheckExists(config.Public.AppPath) {
		folders.Create(config.Public.AppPath)
	}

	if !folders.CheckExists(config.Public.ConfigFile.PluginPath) {
		folders.Create(config.Public.ConfigFile.PluginPath)
	}
}
