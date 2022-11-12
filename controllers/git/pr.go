package git_controller

import (
	"strings"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/app"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/Minnek-Digital-Studio/cominnek/helper"
)

func Pull_request(ticket string, branch string, baseBranch string) (body, title string) {
	bodyFilePath := config.Public.PRBody
	asserts := app.ConfigGlobal.PR.Asserts
	_title := "${branch}"
	variables := []helper.Variables{
		{
			Variable: "${ticket}",
			Value:    ticket,
		},
		{
			Variable: "${branch}",
			Value:    branch,
		},
		{
			Variable: "${base_branch}",
			Value:    baseBranch,
		},
		{
			Variable: "${type}",
			Value:    strings.Split(branch, "/")[0],
		},
	}

	for _, assert := range asserts {
		if assert.Type == "*" {
			bodyFilePath = assert.Body
			_title = assert.Title
		}

		if strings.Contains(branch, assert.Type) {
			bodyFilePath = assert.Body
			_title = assert.Title

			break
		}
	}

	bodyByte := files.Read(bodyFilePath)
	_body := string(bodyByte)

	body = helper.ReplaceValues(_body, variables)
	body = body + `

<small>Created with <a href="https://github.com/Minnek-Digital-Studio/cominnek" target="_blank">Cominnek 🔥</a></small>`

	title = helper.ReplaceValues(_title, variables)

	return
}
