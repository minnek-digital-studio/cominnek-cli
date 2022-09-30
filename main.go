package main

import (
	"github.com/Minnek-Digital-Studio/cominnek/cmd"
	"github.com/Minnek-Digital-Studio/cominnek/config"
)

func main() {
	config.Defaults()
	cmd.Execute()
}
