package controllers

import (
	"fmt"
	"runtime"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/bridge"
	github_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/github"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

var currentVersion = config.Public.Version
var latestVersion = github_controller.GetLatestVersion()
var allOk = true
var osName = runtime.GOOS

func CheckUpdates(printMessage bool) bool {
	if currentVersion != latestVersion {
		if printMessage {
			fmt.Print("\n\n")
			color.HiYellow("🎉🎉🎉 A new version of cominnek is available! 🎉🎉🎉")
			fmt.Println(color.MagentaString(currentVersion), "→ ", color.GreenString(latestVersion))
			fmt.Print("\n")
			fmt.Println("Run", color.HiGreenString("'cominnek update'"), "to update or download the latest version from:")
			color.HiBlue("https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/")
		}

		return true
	}

	return false
}

func installUpdates(route string)  {
	shell.ExecuteCommand(`Start-Process -FilePath "` + route + `" -Argument "/silent" -PassThru`, false)
}

func Update() {
	if !CheckUpdates(false) {
		fmt.Println("🥳🎈 You are using the latest version of cominnek")
		return
	}

	fileName := github_controller.GetLatestFileName()
	url := "https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/download/" + fileName
	route := bridge.DownloadFromURL(url, fileName)
	
	if osName == "windows" {
		installUpdates(route)
	}

	if allOk {
		color.HiBlue("\n🎉🎉🎉 cominnek " + latestVersion + " has been downloaded successfully! 🎉🎉🎉")
	}
}
