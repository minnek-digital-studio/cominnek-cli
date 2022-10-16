package config

import (
	data_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/data"
	"github.com/Minnek-Digital-Studio/cominnek/helper"
)

func Defaults() {
	data_controller.SavePrBody(helper.PRBody, Public.PRBody, false)
}
