package app

import (
	"encoding/json"
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
)

type _assert struct {
	Type       string   `json:"type"`
	Body       string   `json:"body"`
	Title      string   `json:"title"`
	BaseBranch []string `json:"base_branch"`
}

type _PR struct {
	DefaultBranch string    `json:"default_branch"`
	Asserts       []_assert `json:"asserts"`
}

type IConfigLocal struct {
	ProjectKey string `json:"project_key"`
	PR         _PR    `json:"pr"`
}

type IConfigGlobal struct {
	PR _PR `json:"pr"`
}

var ConfigLocal IConfigLocal
var ConfigGlobal IConfigGlobal
var defaultGlb string = `{
  "pr": {
    "asserts": [
      {
        "type": "*",
        "body": "${{default_body}}",
        "title": "${branch}",
        "base_branch": ["develop"]
      },
	  {
		"type": "release",
		"body": "${{default_body}}",
		"title": "${branch} ${base_branch}",
		"base_branch": ["develop", "master"]
	  },
	  {
		"type": "hotfix",
		"body": "${{default_body}}",
		"title": "${branch} ${base_branch}",
		"base_branch": ["develop", "master"]
	  }
    ]
  }
}`

func getFile(file string) []byte {
	if files.CheckExist(file) {
		return files.Read(file)
	}

	return nil
}

func ConfigReader() {
	local := config.Public.ConfigFile.Name
	global := config.Public.ConfigFile.GlobalPath
	var localFile = getFile(local)
	var globalFile = getFile(global)

	if globalFile == nil {
		defaultGlb = strings.ReplaceAll(defaultGlb, "${{default_body}}", config.Public.PRBody)
		files.Create([]byte(defaultGlb), global)
	}

	json.Unmarshal(localFile, &ConfigLocal)
	json.Unmarshal(globalFile, &ConfigGlobal)

	for i, assert := range ConfigGlobal.PR.Asserts {
		println("{")
		println("\tid:", i)
		println("\ttype:", assert.Type)
		println("\ttitle:", assert.Title)
		println("\tbody:", assert.Body)
		println("\tbase_branch:", "[", strings.Join(assert.BaseBranch, ", "), "]")
		println("}")
	}
}
