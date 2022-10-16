package config

import (
	"os"
	"path/filepath"
)

type IConfig struct {
	AppPath   string
	TempPath  string
	KeyPath   string
	TokenPath string
	Version   string
	PRBody    string
}

var userPath, _ = os.UserConfigDir()
var tempPath, _ = os.UserCacheDir()

var cominnekPath = filepath.Join(userPath, ".cominnek")
var cominnekTempPath = filepath.Join(tempPath, ".cominnek")

var Public = IConfig{
	KeyPath:   filepath.Join(cominnekPath, "key.bin"),
	TokenPath: filepath.Join(cominnekPath, "auth.bin"),
	PRBody:    filepath.Join(cominnekPath, "pr-body.md"),
	AppPath:   cominnekPath,
	TempPath:  cominnekTempPath,
	Version:   "v2.0.1",
}
