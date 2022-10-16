package data_controller

import (
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
)

func SavePrBody(data string, file string, overwrite bool) {
	if overwrite {
		files.Delete(file)
	}

	if !files.CheckExist(file) {
		bufer := []byte(data)
		files.Create(bufer, file)
	}
}
