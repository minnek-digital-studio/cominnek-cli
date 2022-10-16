package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	data_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/data"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/files"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var markdownRoute string
var PRConfigCmd = &cobra.Command{
	Use:   "pr",
	Short: "configure the defaults for pull requests",
	Run: func(cmd *cobra.Command, args []string) {
		if !files.CheckExist(markdownRoute) {
			fmt.Println(color.HiRedString("💀Error: ")+ "The file: " + color.HiYellowString(markdownRoute)  +" does not exist.")
			os.Exit(1)
		}

		bodyByte := files.Read(markdownRoute)
		body := string(bodyByte)

		data_controller.SavePrBody(body, config.Public.PRBody, true)
		log.Println(color.HiBlueString("The default pull request body has been set"))
	},
}

func init() {
	PRConfigCmd.PersistentFlags().StringVarP(&markdownRoute, "body", "b", "", "the path to the markdown file to use as the body of the pull request")
}
