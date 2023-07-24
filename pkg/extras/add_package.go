package extras

import (
	"path/filepath"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/bridge"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
)

var url_path = "/minnek-digital-studio/cominnek-src"
var github_repo = "https://github.com" + url_path + "/blob/master"
var github_raw = "https://raw.githubusercontent.com" + url_path + "/master"

func AddPackage(pkg string) bool {
	if strings.Contains(pkg, "flow:") {
		return AddFlow(strings.Split(pkg, ":")[1])
	} else {
		println("🚧 Sorry, package not found")
		return false
	}
}

func AddFlow(flow string) bool {
	url := github_raw + "/flows/" + flow + ".json"
	return download(url, config.Public.FlowPath)
}

func download(url string, path string) bool {
	filename := strings.Split(url, "/")[len(strings.Split(url, "/"))-1]

	if !checkIfEndpointExist(url) {
		println("🚧 Sorry, package not found")
		return false
	}

	downloaded := bridge.DownloadFromURL(url, filename, path)

	if downloaded == "" {
		return false
	}

	return true
}

func checkIfEndpointExist(url string) bool {
	return bridge.CheckIfExist(url)
}

func CheckLocalPackageExist(pkg string) bool {
	filename := filepath.Join(config.Public.FlowPath, pkg+".json")
	return files.CheckExist(filename)
}
