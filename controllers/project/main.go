package project

import (
	"encoding/json"

	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
)

var Config Cominnek

func ReadConfigFile() bool {
	fileNames := []string{
		"mnk-config.json",
	}

	for _, fileName := range fileNames {
		if files.CheckExist(fileName) {
			configByte := files.Read(fileName)
			Config = convertToJSON(configByte).Cominnek
			return true
		}
	}

	return false
}

func convertToJSON(data []byte) Project {
	var project Project
	err := json.Unmarshal([]byte(data), &project)

	if err != nil {
		panic(err)
	}

	return project
}

func GetConfigByName(name string) Branch {
	for _, branch := range Config.Git.Branches {
		if branch.Name == name {
			return branch
		}
	}

	return Branch{}
}
