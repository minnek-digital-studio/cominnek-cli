package app

import (
	"path/filepath"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/folders"
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
	Name    string     `json:"name"`
	Help    string     `json:"help"`
	Exec    string     `json:"exec"`
	Flags   []IFlag    `json:"flags"`
	Scripts []IScripts `json:"scripts"`
}

type IPlugin struct {
	Version string     `json:"version"`
	Process string     `json:"process"`
	Name    string     `json:"name"`
	Help    string     `json:"help"`
	Exec    string     `json:"exec"`
	Flags   []IFlag    `json:"flags"`
	Scripts []IScripts `json:"scripts"`
}

type IConfigGlobal struct {
	PR      _PR      `json:"pr"`
	Plugins []string `json:"plugins"`
}

type IConfigLocal struct {
	PR _PR `json:"pr"`
}

var ConfigLocal IConfigLocal
var ConfigGlobal IConfigGlobal
var PluginsConfig []IPlugin
var fileName = "index.cmk"

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

func pluginScriptsExec(pluginConfig []IScripts, pluginPath string, process string) {
	var cmd string

	for i, script := range pluginConfig {
		if pluginConfig[i].Exec != "" {
			cmd = `"` + filepath.Join(pluginPath, script.Exec) + `"`

			if process != "" {
				cmd = process + " " + cmd
			}

			pluginConfig[i].Exec = cmd

			if len(script.Scripts) > 0 {
				pluginScriptsExec(script.Scripts, pluginPath, process)
			}
		}
	}
}

func commandProcessor(pluginPath string, pluginConfig IPlugin) {
	cmd := `"` + filepath.Join(pluginPath, pluginConfig.Exec) + `"`

	if pluginConfig.Process != "" {
		cmd = pluginConfig.Process + " " + cmd
	}

	pluginConfig.Exec = cmd

	pluginScriptsExec(pluginConfig.Scripts, pluginPath, pluginConfig.Process)
	PluginsConfig = append(PluginsConfig, pluginConfig)
}

func getPlugin(plugin string) string {
	pluginPath := filepath.Join(config.Public.ConfigFile.PluginPath, plugin)

	if !folders.CheckExists(pluginPath) || !files.CheckExist(filepath.Join(pluginPath, fileName)) {
		println("Plugin " + plugin + " not found")
		return ""
	}

	return pluginPath
}

func pluginsReader() {

	for _, plugin := range ConfigGlobal.Plugins {
		pluginViper := viper.New()

		pluginPath := getPlugin(plugin)

		if pluginPath == "" {
			continue
		}

		pluginViper.SetConfigName(fileName)
		pluginViper.SetConfigType("json")
		pluginViper.AddConfigPath(pluginPath)

		err := pluginViper.ReadInConfig()

		if err != nil {
			panic(err)
		}

		var pluginConfig IPlugin
		pluginViper.Unmarshal(&pluginConfig)

		commandProcessor(pluginPath, pluginConfig)
	}
}

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
	pluginsReader()
}
