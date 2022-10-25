package pkg_action

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/logger"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/fatih/color"
)

func getLists(unstaged, list []string) (defaults []string, listUnstaged []string) {
	if len(unstaged) == len(list) {
		listUnstaged = list
		logger.Warning("All files are unstaged")
		return
	}

	if len(unstaged) > 0 {
		for _, item := range list {
			match := ""
			logger.PrintLn("{")
			logger.PrintLn("\titem: ", item)
			logger.PrintLn("\tCheking: [")
			for i, unstagedItem := range unstaged {
				if unstagedItem == "" {
					continue
				}

				logger.PrintLn("\t\t" + unstagedItem)

				if strings.Contains(item, unstagedItem) {
					match = unstagedItem
					unstaged = append(unstaged[:i], unstaged[i+1:]...)
					break
				}
			}
			logger.PrintLn("\t]")

			if match != "" {
				listUnstaged = append(listUnstaged, item)
				logger.PrintLn("\tUntraked: ", color.HiGreenString("Yes"))
				logger.PrintLn("}")
				continue
			}

			logger.PrintLn("\tUntraked: ", color.HiRedString("No"))
			defaults = append(defaults, item)
			logger.PrintLn("}")
		}
	} else {
		defaults = list
		logger.Warning("All files are staged")
	}

	return
}

func addToStage(raw []string) {
	loading.Start("Adding files to stage ")
	files := config.AppData.Commit.Files
	filesLen := len(files)
	rawLen := len(raw)

	defer loading.Stop()
	defer logger.Success("Successfully staged " + strconv.Itoa(filesLen) + " files")

	if filesLen == rawLen {
		git_controller.AddAll()
		return
	}

	if filesLen > 0 {
		git_controller.Reset()
	}

	for _, _file := range raw {
		if filesLen == 0 {
			break
		}

		for i, file := range files {
			if strings.Contains(file, _file) {
				git_controller.AddSpecific(filepath.Join("./", _file))
				files = append(files[:i], files[i+1:]...)
				break
			}
		}
	}
}

func processFiles(raw []string, unstaged []string, list []string) (newList []string, changesMsg string, defaults []string) {
	if len(raw) == 0 {
		println("No changes to commit ✅")
		os.Exit(0)
		return
	}

	defaults, listUnstaged := getLists(unstaged, list)

	newList = func() []string {
		if len(defaults) > len(listUnstaged) {
			return append(listUnstaged, defaults...)
		} else {
			return append(defaults, listUnstaged...)
		}
	}()

	changesMsg = func() string {
		lenDefaults := len(defaults)
		lenListUnstaged := len(listUnstaged)

		coloredLenDefaults := color.HiRedString(strconv.Itoa(lenDefaults))
		coloredLenListUnstaged := color.HiGreenString(strconv.Itoa(lenListUnstaged))

		if lenDefaults > 0 {
			return "(Changes to commit: " + coloredLenDefaults + " | Changes to stage: " + coloredLenListUnstaged + ")"
		}

		return "(Changes to commit: " + coloredLenListUnstaged + ")"
	}()

	return
}

func commitQuestions(list, raw []string) {
	if !config.AppData.Commit.AddAll {
		unstaged := git_controller.ListUnstageChanges()
		newList, changesMsg, defaults := processFiles(raw, unstaged, list)

		loading.Stop()

		if len(config.AppData.Commit.Files) < 1 {
			ask.One(&survey.MultiSelect{
				Message:       "Select files to commit " + changesMsg + ":",
				Options:       newList,
				Help:          "Use space to select files, enter to continue",
				FilterMessage: "Type to filter files",
				Default:       defaults,
			}, &config.AppData.Commit.Files, survey.WithValidator(survey.Required))
		}
	}

	if config.AppData.Commit.Type == "" {
		ask.One(&survey.Select{
			Message:       "Select Type :",
			Options:       config.Public.Commits.Types,
			Help:          "Use space to select files, enter to continue",
			FilterMessage: "Type to filter files",
		}, &config.AppData.Commit.Type, survey.WithValidator(survey.Required))
	}

	if config.AppData.Commit.Scope == "" {
		ask.One(&survey.Input{
			Message: "Enter scope :",
		}, &config.AppData.Commit.Scope)
	}

	if config.AppData.Commit.Message == "" {
		ask.One(&survey.Input{
			Message: "Commit message:",
			Help:    "Enter commit message",
		}, &config.AppData.Commit.Message, survey.WithValidator(survey.Required))
	}

	if config.AppData.Commit.Body == "" {
		ask.One(&survey.Multiline{
			Message: "Commit body:",
			Help:    "Enter commit body",
		}, &config.AppData.Commit.Body)
	}
}

func Commit() {
	raw := []string{}
	list := []string{}

	if !config.AppData.Commit.AddAll {
		loading.Start("Checking files status ")
		list, raw = git_controller.ListChanges()
	}

	commitQuestions(list, raw)

	// helper.PrintName()

	msg := config.AppData.Commit.Message
	body := config.AppData.Commit.Body
	ctype := config.AppData.Commit.Type
	scope := config.AppData.Commit.Scope

	if !config.AppData.Commit.AddAll {
		addToStage(raw)
	} else {
		git_controller.AddAll()
	}

	git.Commit(msg, body, ctype, scope)
}
