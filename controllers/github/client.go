package github_controller

import (
	"context"
	"fmt"
	"os"

	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/security"
	"github.com/fatih/color"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

var ctx context.Context = context.Background();

func client()* github.Client {
	if (!files.CheckExist(FileName) ) { 
		loading.Stop()
		fmt.Println("No token found, please login. run", color.HiGreenString("cominnek auth login"))
		os.Exit(1)
		return nil
	}

	bite := files.Read(FileName)
	code := security.Decrypt(bite, security.GetKey())

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: code},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client
}
