package controllers

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/bridge"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	github_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/github"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
	"github.com/hashicorp/go-version"
)

var currentVersion = config.Public.Version
var allOk = true
var osName = runtime.GOOS
var maxToCheck = 10
var showUpdateMessage bool
var latestVersion string

func CheckUpdates() bool {
	if !Connected() {
		return false
	}

	latestVersion = github_controller.GetLatestVersion()

	_current, _ := version.NewVersion(currentVersion)
	_latest, _ := version.NewVersion(latestVersion)

	showUpdateMessage = _current.LessThan(_latest)

	return showUpdateMessage
}

func PrintUpdateMessage() {
	if showUpdateMessage {
		fmt.Print("\n\n")
		color.HiYellow("🎉🎉🎉 A new version of cominnek is available! 🎉🎉🎉")
		fmt.Println(color.MagentaString(currentVersion), "→ ", color.GreenString(latestVersion))
		fmt.Print("\n")
		fmt.Println("Run", color.HiGreenString("'cominnek update'"), "to update or download the latest version from:")
		color.HiBlue("https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/")
	}
}

func checkDistToUnmount(mountOut string, firstNumber int, lastNumber int) string {
	str := "/dev/disk" + fmt.Sprint(firstNumber) + "s" + fmt.Sprint(lastNumber)
	disk := strings.Contains(mountOut, str)

	if disk {
		return str
	}

	if lastNumber == 5 && firstNumber < maxToCheck {
		return checkDistToUnmount(mountOut, firstNumber+1, 1)
	}

	if lastNumber == 5 && firstNumber == maxToCheck {
		return ""
	}

	return checkDistToUnmount(mountOut, firstNumber, lastNumber+1)
}

func getMountedDisc(mountOut string, name string, num int) string {
	str := "/Volumes/" + name
	var preTest = false

	if num == 0 {
		preTest = strings.Contains(mountOut, str+" ")
	}

	if num == maxToCheck {
		return str
	}

	if num > 0 {
		str = str + " " + fmt.Sprint(num)
	}

	mounted := strings.Contains(mountOut, str)

	if mounted && !preTest {
		return str
	}

	return getMountedDisc(mountOut, name, num+1)
}

func mountDisk(route string, name string) (string, string) {
	out, _, err := shell.Out("hdiutil attach " + route)

	if err != nil {
		fmt.Println(err)
		allOk = false
	}

	disk := checkDistToUnmount(out, 1, 1)
	mounted := getMountedDisc(out, name, 0)

	return disk, mounted
}

func checkUpdated(latestVersion string) {
	loading.Start("🔎 Checking if cominnek has been updated")
	out, _, err := shell.Out("cominnek -v")

	if err != nil {
		loading.Stop()
		fmt.Println(err)
		os.Exit(1)
	}

	loading.Stop()
	if strings.Contains(out, latestVersion) {
		fmt.Println("🥳🎈 cominnek was successfully updated")
	} else {
		fmt.Println("🤔 Something went wrong, try to update again")
	}
}

func installUpdates(route string, fileName string) {
	latestVersion := github_controller.GetLatestVersion()

	if osName == "windows" {
		shell.ExecuteCommand(`Start-Process -FilePath "`+route+`" -Argument "/silent" -PassThru`, false)

		if allOk {
			color.HiBlue("\n🎉🎉🎉 cominnek " + latestVersion + " has been downloaded successfully! 🎉🎉🎉")
		}
	}

	if osName == "darwin" {
		name := strings.Split(fileName, ".dmg")[0]
		loading.Start("📦 Installing " + color.HiGreenString(name))

		disk, mounted := mountDisk(route, name)
		shell.ExecuteCommand("cd "+mounted+"; bash installer.sh", false)
		shell.ExecuteCommand("hdiutil detach "+disk, false)

		loading.Stop()

		checkUpdated(latestVersion)

		files.Delete(route)
	}

	if osName == "linux" {
		shell.ExecuteCommand("sudo dpkg -i "+route, false)

		checkUpdated(latestVersion)
		files.Delete(route)
	}
}

func Update() {
	if !Connected() {
		fmt.Println("🤔 You are not connected to the internet")
		return
	}

	if !CheckUpdates() {
		fmt.Println("🥳🎈 You are using the latest version of cominnek")
		return
	}

	fileName := github_controller.GetLatestFileName()
	url := "https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/download/" + fileName
	route := bridge.DownloadTempFromURL(url, fileName)

	installUpdates(route, fileName)
}
