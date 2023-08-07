package extras

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/folders"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/helper"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/fatih/color"
)

var url = "https://api.github.com/repos/minnek-digital-studio/cominnek-src/git/trees/master?recursive=1"
var configTemplate = `
{
  "cominnek": {
    "git": {
      "flow": "${flow}"
    }
  }
}
`

func InitProject() {
	flows := checkFromCache()
	var flow string

	if CheckIfConfigExists() {
		var overwrite bool

		ask.One(&survey.Confirm{
			Message: "Config already exists, do you want to overwrite it?",
		}, &overwrite, nil)

		if overwrite == false {
			return
		}
	}

	ask.One(&survey.Select{
		Message: "Which flow do you want to use?",
		Options: flows,
	}, &flow, nil)

	saveConfig(flow)

	helper.PrintName()
	color.Green("Init project successfully")
}

func CheckIfConfigExists() bool {
	for _, fileName := range config.Public.ConfigFilesNames {
		if files.CheckExist(fileName) {
			return true
		}
	}

	return false
}

func saveConfig(flow string) {
	if CheckLocalPackageExist(flow) == false {
		println("Downloading " + flow)

		if !AddPackage("flow:" + flow) {
			println("\n" + color.RedString("Something goes wrong downloading the flow"))
			os.Exit(1)
		}
	}

	configJson := []byte(strings.Replace(configTemplate, "${flow}", flow, 1))
	err := ioutil.WriteFile(config.Public.ConfigFilesNames[0], configJson, 0644)

	if err != nil {
		panic(err)
	}
}

func getFlowsFromRepo() []string {
	loading.Start("Fetching flows from repo")

	defer loading.Stop()
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var ghResponse GitHubResponse

	json.Unmarshal(body, &ghResponse)
	filenames := []string{}
	for _, item := range ghResponse.Tree {
		if item.Type == "blob" {
			if strings.Contains(item.Path, "flows/") && strings.Contains(item.Path, ".json") {
				filename := strings.Replace(item.Path, "flows/", "", 1)
				filename = strings.Replace(filename, ".json", "", 1)
				filenames = append(filenames, filename)
			}
		}
	}

	return filenames
}

func checkFromCache() []string {
	currentDate := time.Now().Unix()

	if !folders.CheckExists(config.Public.TempPath) {
		folders.Create(config.Public.TempPath)
	}

	if files.CheckExist(config.Public.CacheFile) == false {
		return updateCache()
	}

	cacheFile, err := ioutil.ReadFile(config.Public.CacheFile)

	if err != nil {
		panic(err)
	}

	var cache CacheFile

	json.Unmarshal(cacheFile, &cache)

	// check if cache is older than 30 Minutes
	if currentDate-cache.FlowCache.LastUpdate > 30*60 {
		return updateCache()
	}

	return cache.FlowCache.Flows
}

func updateCache() []string {
	flows := getFlowsFromRepo()

	cache := CacheFile{
		FlowCache: FlowCache{
			Flows:      flows,
			LastUpdate: time.Now().Unix(),
		},
	}

	cacheJson, err := json.Marshal(cache)

	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(config.Public.CacheFile, cacheJson, 0644)

	if err != nil {
		panic(err)
	}

	return flows
}
