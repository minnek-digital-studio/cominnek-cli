package git_controller

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/Minnek-Digital-Studio/cominnek/helper"
)

func Pull_request(ticket string, branch string) string {
	variables := []helper.Variables{
		{
			Variable: "${ticket}",
			Value:    ticket,
		},
		{
			Variable: "${branch}",
			Value:    branch,
		},
	}

	bodyByte := files.Read(config.Public.PRBody)
	body := string(bodyByte)

	msg := helper.ReplaceValues(body, variables)

	return msg
}
