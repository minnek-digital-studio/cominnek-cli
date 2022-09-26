package github_controller

import "fmt"

func TestConnection() {
	client := client()
	user, _, err := client.Users.Get(ctx, "")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user.GetLogin())
}