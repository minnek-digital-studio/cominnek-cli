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
		Merge string
	}
	Publish struct {
		Ticket string
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
}

type IConfig struct {
	AppPath   string
	TempPath  string
	KeyPath   string
	TokenPath string
	Version   string
	Commits   ICommit
	PRBody    string
	Logs      bool
}

var userPath, _ = os.UserConfigDir()
var tempPath, _ = os.UserCacheDir()

var cominnekPath = filepath.Join(userPath, ".cominnek")
var cominnekTempPath = filepath.Join(tempPath, ".cominnek")

var Public = IConfig{
	Version:   "v2.2.0",
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
