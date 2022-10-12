package github_controller

import (
	"fmt"

	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
)

func TestConnection() {
	loading.Start("Testing connection to Github ")
	client := client()
	user, _, err := client.Users.Get(ctx, "")

	if err != nil {
		fmt.Println(err)
		return
	}

	loading.Stop()
	fmt.Println(user.GetHTMLURL())
	fmt.Println("Connected to Github as", user.GetLogin())
}
