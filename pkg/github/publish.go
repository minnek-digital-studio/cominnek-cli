package github

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
)

func Publish(msg string, body string, ctype string,  scope string) {
	if(msg != "") {
		git.Add();
		git.Status();
		git.Commit(msg, body, ctype, scope);
	}
	git.PushPublish();
	CreatePullRequest();
}