package project

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/folders"
	project_structs "github.com/Minnek-Digital-Studio/cominnek/controllers/project/structs"
	"github.com/fatih/color"
)

var Config project_structs.Cominnek
var exitOnFail bool

func ReadConfigFile(_exitOnFail bool) bool {
	exitOnFail = _exitOnFail
	fileNames := config.Public.ConfigFilesNames

	for _, fileName := range fileNames {
		if files.CheckExist(fileName) {
			configByte := files.Read(fileName)
			converted := convertToJSON(configByte).Cominnek
			Config = getFlow(converted)
			return true
		}
	}

	return false
}

func convertToJSON(data []byte) project_structs.Project {
	var project project_structs.Project
	err := json.Unmarshal([]byte(data), &project)

	if err != nil {
		panic(err)
	}

	return project
}

func GetConfigByName(name string) project_structs.Branch {
	for _, branch := range Config.Git.Branches {
		if branch.Name == name {
			return branch
		}
	}

	return project_structs.Branch{}
}

func getFlow(cmk project_structs.Cominnek) project_structs.Cominnek {
	flowStr := cmk.Git.Flow

	if flowStr == "custom" {
		return cmk
	}

	if !folders.CheckExists(config.Public.FlowPath) {
		folders.Create(config.Public.FlowPath)
	}

	var flow project_structs.Git
	fileName := filepath.Join(config.Public.FlowPath, flowStr+".json")

	if !files.CheckExist(fileName) {
		println("Sorry but the flow " + flowStr + " does not exist")
		println("Try to download it with the command:")

		color.Green("\tcominnek add flow:" + flowStr)

		if exitOnFail {
			os.Exit(1)
		} else {
			return cmk
		}
	}

	file := files.Read(fileName)

	err := json.Unmarshal([]byte(file), &flow)

	if err != nil {
		panic(err)
	}

	cmk.Git = flow

	return cmk
}
