package pkg_action

import (
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/fatih/color"
)

func getLists(unstaged, list []string, log bool) (defaults []string, listUnstaged []string) {
	printLog := func(args ...string) {
		if log {
			for i, arg := range args {
				addSpace := func() string {
					if i > 0 && i < len(args)-1 {
						return " "
					}

					return ""
				}
				print(arg, addSpace())
			}
			println("")
		}
	}

	if len(unstaged) == len(list) {
		listUnstaged = list
		return
	}

	if len(unstaged) > 0 {
		for _, item := range list {
			match := ""
			printLog("{")
			printLog("\titem: ", item)
			printLog("\tCheking: [")
			for i, unstagedItem := range unstaged {
				if unstagedItem == "" {
					continue
				}

				printLog("\t\t" + unstagedItem)

				if strings.Contains(item, unstagedItem) {
					match = unstagedItem
					unstaged = append(unstaged[:i], unstaged[i+1:]...)
					break
				}
			}
			printLog("\t]")

			if match != "" {
				listUnstaged = append(listUnstaged, item)
				printLog("\tMatch: ", color.HiGreenString("Yes"))
				continue
			}

			printLog("\tMatch: ", color.HiRedString("No"), item)
			defaults = append(defaults, item)
			printLog("}")
		}
	} else {
		listUnstaged = list
	}

	return
}

func Commit() {
	list, raw := git_controller.ListChanges()
	unstaged := git_controller.ListUnstageChanges()

	if len(raw) == 0 {
		println("No changes to commit ✅")
		return
	}

	defaults, listUnstaged := getLists(unstaged, list, false)

	getList := func() []string {
		if len(defaults) > len(listUnstaged) {
			return append(listUnstaged, defaults...)
		} else {
			return append(defaults, listUnstaged...)
		}
	}

	countChangesMsg := func() string {
		_list, _ := git_controller.ListChanges()
		_unstaged := git_controller.ListUnstageChanges()
		_defaults, _listUnstaged := getLists(_unstaged, _list, false)
		lenDefaults := len(_defaults)
		lenListUnstaged := len(_listUnstaged)

		coloredLenDefaults := color.HiRedString(strconv.Itoa(lenDefaults))
		coloredLenListUnstaged := color.HiGreenString(strconv.Itoa(lenListUnstaged))

		if lenDefaults > 0 {
			return "(Changes to commit: " + coloredLenDefaults + " | Changes to stage: " + coloredLenListUnstaged + ")"
		}

		return "(Changes to commit: " + coloredLenListUnstaged + ")"
	}

	survey.AskOne(&survey.MultiSelect{
		Message:       "Select files to commit " + countChangesMsg() + ":",
		Options:       getList(),
		Help:          "Use space to select files, enter to continue",
		FilterMessage: "Type to filter files",
		Default:       defaults,
	}, &config.AppData.Commit.Files, survey.WithValidator(survey.Required))
}
