package github

import (
	"log"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
)

func Publish(msg string, body string, ctype string,  scope string, ticket string) {
	if(msg != "") {
		git.Add();
		git.Status();
		git.Commit(msg, body, ctype, scope);
	}
	git.PushPublish();
	CreatePullRequest(ticket);
	log.Println("Publish complete")
}