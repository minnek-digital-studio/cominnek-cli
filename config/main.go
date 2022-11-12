package config

import (
	"os"
	"path/filepath"
)

type ICommit struct {
	Types []string
}

type IAppData struct {
	Action string
	Commit struct {
		Message string
		Files   []string
		AddAll  bool
		Body    string
		Type    string
		Scope   string
	}
	Push struct {
		Merge        string
		IgnoreCommit bool
	}
	Publish struct {
		Ticket       string
		IgnoreCommit bool
	}
	PullRequest struct {
		Ticket string
		Base   string
	}
	Flow struct {
		Type   string
		Ticket string
		Stash  bool
	}
	Stash struct {
		Branch string
	}
	Merge struct {
		Branch string
	}
}

type IConfigFile struct {
	Name       string
	GlobalPath string
	UserPath   string
	PluginPath string
}
type IConfig struct {
	AppPath    string
	TempPath   string
	KeyPath    string
	TokenPath  string
	Version    string
	Commits    ICommit
	PRBody     string
	Logs       bool
	ConfigFile IConfigFile
	HomePath   string
}

var userPath, _ = os.UserConfigDir()
var tempPath, _ = os.UserCacheDir()
var homePath, _ = os.UserHomeDir()

var cominnekPath = filepath.Join(userPath, ".cominnek")
var cominnekTempPath = filepath.Join(tempPath, ".cominnek")
var configFileName = ".cominnekrc"
var Public = IConfig{
	Version: "v2.4.0",
	ConfigFile: IConfigFile{
		Name:       configFileName,
		GlobalPath: filepath.Join(cominnekPath, configFileName),
		UserPath:   filepath.Join(homePath, configFileName),
		PluginPath: filepath.Join(cominnekPath, "plugins"),
	},
	HomePath:  homePath,
	KeyPath:   filepath.Join(cominnekPath, "key.bin"),
	TokenPath: filepath.Join(cominnekPath, "auth.bin"),
	PRBody:    filepath.Join(cominnekPath, "pr-body.md"),
	AppPath:   cominnekPath,
	TempPath:  cominnekTempPath,
	Logs:      false,
	Commits: ICommit{
		Types: []string{
			"feat",
			"fix",
			"docs",
			"style",
			"refactor",
			"perf",
			"test",
			"build",
			"ci",
			"chore",
			"revert",
		},
	},
}

var AppData = IAppData{}
