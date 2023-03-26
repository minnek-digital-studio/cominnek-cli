package github_controller

import (
	"fmt"
	"os"

	"github.com/google/go-github/v47/github"
)

func GetCurrentUser()*github.User {
	client := client();
	user, _, err := client.Users.Get(ctx, "")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return user
}
