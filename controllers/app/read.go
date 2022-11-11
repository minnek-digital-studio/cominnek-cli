package app

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/spf13/viper"
)

type _assert struct {
	Type       string   `json:"type"`
	Body       string   `json:"body"`
	Title      string   `json:"title"`
	BaseBranch []string `json:"base_branch"`
}

type _PR struct {
	Asserts []_assert `json:"asserts"`
}

type IFlag struct {
	Short    string      `json:"short"`
	Long     string      `json:"long"`
	Type     string      `json:"type"`
	Help     string      `json:"help"`
	Default  interface{} `json:"default"`
	Required bool        `json:"required"`
}

type IScripts struct {
	Name    string     `json:"command"`
	Help    string     `json:"help"`
	Exec    string     `json:"exec"`
	Flags   []IFlag    `json:"flags"`
	Scripts []IScripts `json:"scripts"`
}

type IPlugin struct {
	Name    string     `json:"name"`
	Path    string     `json:"path"`
	Help    string     `json:"help"`
	Version string     `json:"version"`
	Flags   []IFlag    `json:"flags"`
	Scripts []IScripts `json:"scripts"`
}

type IConfigGlobal struct {
	PR      _PR       `json:"pr"`
	Plugins []IPlugin `json:"plugins"`
}

type IConfigLocal struct {
	PR _PR `json:"pr"`
}

var ConfigLocal IConfigLocal
var ConfigGlobal IConfigGlobal
var defaultGlb string = `{
  "plugins": [],
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

func ConfigReader() {
	global := config.Public.ConfigFile.GlobalPath

	if !files.CheckExist(global) {
		files.Create([]byte(defaultGlb), global)
	}

	viper.SetConfigName(".cominnekrc")
	viper.SetConfigType("json")
	viper.AddConfigPath(config.Public.HomePath)
	viper.AddConfigPath(config.Public.AppPath)

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	viper.Unmarshal(&ConfigGlobal)
}
