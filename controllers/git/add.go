package git_controller

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
	"github.com/fatih/color"
)

func Add() string {
	return "git add ."
}

func AddSpecific(file string) {
	cmd := "git add -A -- " + file

	_, errOut, err := shell.Out(cmd)

	if err != nil {
		println(color.HiRedString("Error: ") + errOut)
	}
}

func AddAll() {
	cmd := Add()

	_, errOut, err := shell.Out(cmd)

	if err != nil {
		println(color.HiRedString("Error: ") + errOut)
	}
}
