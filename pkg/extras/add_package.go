package extras

import (
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/bridge"
)

var url_path = "/minnek-digital-studio/cominnek-src"
var github_repo = "https://github.com" + url_path + "/blob/master"
var github_raw = "https://raw.githubusercontent.com" + url_path + "/master"

func AddPackage(pkg string) {
	if strings.Contains(pkg, "flow:") {
		AddFlow(strings.Split(pkg, ":")[1])
	} else {
		println("🚧 Sorry, package not found")
	}

}

func AddFlow(flow string) {
	url := github_raw + "/flows/" + flow + ".json"
	download(url, config.Public.FlowPath)
}

func download(url string, path string) {
	filename := strings.Split(url, "/")[len(strings.Split(url, "/"))-1]

	if !checkIfExist(url) {
		println("🚧 Sorry, package not found")
		return
	}

	bridge.DownloadFromURL(url, filename, path)
}

func checkIfExist(url string) bool {
	return bridge.CheckIfExist(url)
}
