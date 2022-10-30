package controllers

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
)

func Confirm(msg string, defaultAnswer bool) bool {
	response := defaultAnswer

	ask.One(&survey.Confirm{
		Message: msg,
		Default: false,
	}, &response)

	return response
}
