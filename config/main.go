package config

import (
	"os"
	"path/filepath"
)

type IConfig struct {
	AppPath   string
	KeyPath   string
	TokenPath string
	Version   string
	PRBody    string
}

var userPath, _ = os.UserConfigDir()
var cominnekPath = filepath.Join(userPath, ".cominnek")

var Public = IConfig{
	KeyPath:   filepath.Join(cominnekPath, "key.bin"),
	TokenPath: filepath.Join(cominnekPath, "auth.bin"),
	PRBody:    filepath.Join(cominnekPath, "pr-body.md"),
	AppPath:   cominnekPath,
	Version:   "2.0.1",
}
