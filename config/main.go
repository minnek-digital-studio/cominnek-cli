package config

import (
	"os"
	"path/filepath"
	"time"

	project_structs "github.com/Minnek-Digital-Studio/cominnek/controllers/project/structs"
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
		Data   project_structs.Branch
		Stash  bool
		Ticket string
	}
	Stash struct {
		Branch string
	}
	Merge struct {
		Branch string
	}
	Clone struct {
		Url string
	}
	CreateRepo struct {
		Name string
	}
	Reset struct {
		Type    string
		Commit  string
		Number  string
		Target  string
		Confirm bool
	}
}

type IConfig struct {
	AppPath   string
	TempPath  string
	FlowPath  string
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
	Version:   "v4.0.0-beta",
	KeyPath:   filepath.Join(cominnekPath, "key.bin"),
	TokenPath: filepath.Join(cominnekPath, "auth.bin"),
	PRBody:    filepath.Join(cominnekPath, "pr-body.md"),
	FlowPath:  filepath.Join(cominnekPath, "flows"),
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
