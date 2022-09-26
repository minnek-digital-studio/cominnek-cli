package github_controller

import (
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/security"
	"github.com/cli/oauth"
)

var FileName string = config.Public.TokenPath

func Login() {
	flow := &oauth.Flow{
		Host:     oauth.GitHubHost("https://github.com"),
		ClientID: config.Private.GithubClient,	// this come from config.priv.go
		Scopes:   []string{"repo", "read:org", "repo:status", "user", "gist", "project"},
	}

	accessToken, err := flow.DetectFlow()
	if err != nil {
		panic(err)
	}


	encrypted := security.Encrypt(accessToken.Token, security.GetKey())
	files.Create(encrypted, FileName)
}
