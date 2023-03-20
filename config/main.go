package config

import (
	"os"
	"path/filepath"
	"time"
)

type ICommit struct {
	Types []string
}

type IAppData struct {
	Action string
	Start  time.Time
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
	Branch struct {
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
	Reset struct {
		Type string
		Commit string
		Number string
		Target string
		Confirm bool
	}
	Clone struct {
		Url string
	}
	CreateRepo struct {
		Name string
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
	Version:   "v3.0.0-alpha.3",
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
