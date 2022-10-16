package github_controller

import (
	"fmt"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/fatih/color"
)

func Logout() {
	loading.Start("Logging out of Github ")
	client := client()
	user, _, err := client.Users.Get(ctx, "")
	name := user.GetName();
	username := user.GetLogin();

	loading.Stop()
	loading.Start("Deleting Github credentials for " + username)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	files.Delete(FileName)
	files.Delete(config.Public.KeyPath)

	if !files.CheckExist(FileName) {
		loading.Stop()
		fmt.Println(color.HiGreenString(username), "You have been logged out of Github 🔒")
		fmt.Println("You can login again with the command", color.HiGreenString("cominnek auth login"))
		fmt.Println("\nGoodbye", color.HiGreenString(name), "👋")
	}
}