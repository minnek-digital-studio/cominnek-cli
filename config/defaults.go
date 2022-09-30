package config

import (
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/Minnek-Digital-Studio/cominnek/helper"
)

func Defaults() {
	if !files.CheckExist(Public.PRBody) {
		bufer := []byte(helper.PRBody)
		files.Create(bufer, Public.PRBody)
	}
}
