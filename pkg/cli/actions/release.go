package pkg_action

import (
	"fmt"
	"os"
	"time"

	"github.com/AlecAivazis/survey/v2"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	github_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/github"
	"github.com/Minnek-Digital-Studio/cominnek/helper"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
	"github.com/fatih/color"
	"github.com/hashicorp/go-version"
)

var _continue bool
var committed = false

func Release() {
	println(color.HiRedString("IMPORTANT") + ": This is a " + color.HiYellowString("beta") + " feature, we recommend you to be careful when using it")
	println("You can report any issue at " + color.HiBlueString(github_controller.GetRepoURL("Minnek-Digital-Studio", "cominnek")+"/issues"))
	print("\n")
	ask.One(&survey.Confirm{
		Message: "Do you want to continue?",
		Default: false,
	}, &_continue)

	if !_continue {
		os.Exit(0)
	}

	checkChanges()

	events.App.On("cleanup", func(payload ...interface{}) {
		color.Yellow("\nResetting\n")
		git.ResetAll()
		color.HiGreen("Done")

		if committed {
			git_controller.ResetBy("hard", "1", "")
		}

		os.Exit(0)
	})

	version := generateChangelog()

	releaseQuestion(version)
}

func releaseQuestion(version string) {
	var release bool

	ask.One(&survey.Confirm{
		Message: fmt.Sprintf("Do you want to release version %s?", version),
		Default: false,
	}, &release)

	if !release {
		events.App.Emit("cleanup")
	}

	git.Add()
	git.Commit("release v"+version, "", "chore", "")

	committed = true

	git.AddTag(version)
	git.PublishTag(version)

	origin := git_controller.GetOrigin()
	url := github_controller.GetRepoURL(origin.Owner, origin.Repo)

	fmt.Println("Pushing to", url)
	git.Push()
}

func generateChangelog() string {
	last := git.LastTag()

	currentVersion, err := version.NewVersion(last)

	if err != nil {
		currentVersion, _ = version.NewVersion("0.0.0")
	}

	origin := git_controller.GetOrigin()
	url := github_controller.GetRepoURL(origin.Owner, origin.Repo)

	commits := git.GetAllCommitsForRelease(last)
	oldChangelog := getReleaseFileData()
	var changelog string
	var features string
	var fixes string
	var performance string
	allowedTypes := []string{"feat", "fix", "perf"}
	currentDate := time.Now().Format("2006-01-02")

	for _, commit := range commits {
		if len(commit) > 0 {
			data := getData(commit)
			if data.Type == "" {
				continue
			}

			if !isInArray(data.Type, allowedTypes) {
				continue
			}

			msg := fmt.Sprintf("- %s ([%s](%s))\n", data.Message, data.Hash, url+"/commit/"+data.Hash)

			if data.Scope != "" {
				msg = fmt.Sprintf("- **%s:** %s ([%s](%s))\n", data.Scope, data.Message, data.Hash, url+"/commit/"+data.Hash)
			}

			if data.Type == "feat" {
				features += msg
			}

			if data.Type == "fix" {
				fixes += msg
			}

			if data.Type == "perf" {
				performance += msg
			}
		}
	}

	if len(features) == 0 && len(fixes) == 0 && len(performance) == 0 {
		println("No changes found")
		os.Exit(0)
	}

	var _version string

	if len(features) > 0 {
		_version = updateVersion(currentVersion.String(), "minor")
	} else if len(fixes) > 0 {
		_version = updateVersion(currentVersion.String(), "patch")
	} else if len(performance) > 0 {
		_version = updateVersion(currentVersion.String(), "patch")
	}

	changesURL := url + "/compare/" + last + "..." + _version
	changelog += fmt.Sprintf("## [%s](%s) (%s)\n", _version, changesURL, currentDate)

	if len(features) > 0 {
		changelog += "\n### Features\n"
		changelog += features
	}

	if len(fixes) > 0 {
		changelog += "\n### Bug Fixes\n"
		changelog += fixes
	}

	if len(performance) > 0 {
		changelog += "\n### Performance\n"
		changelog += performance
	}

	println("Changelog preview:\n")
	fmt.Println(changelog)

	var confirm bool

	ask.One(&survey.Confirm{
		Message: "Do you want to create the release?",
		Default: false,
	}, &confirm)

	if !confirm {
		events.App.Emit("cleanup")
	}

	changelog += "\n" + oldChangelog
	saveChangelog(changelog)

	println("\n")
	return _version
}

func checkChanges() {
	changes, _ := git_controller.ListChanges()
	currentBranch := git_controller.GetCurrentBranch()

	if len(changes) > 0 {
		println("You have changes that are not committed")
		println("Please commit before release")

		println("Run " + color.HiGreenString("'cominnek commit'") + " to commit your changes")
		var commit_it bool
		ask.One(&survey.Confirm{
			Message: "Do you want to run 'cominnek commit'?",
			Default: false,
		}, &commit_it)

		if !commit_it {
			os.Exit(1)
		}

		helper.PrintName()
		Commit(true)

		println("\n\nReady to release, checking changes\n")
		checkChanges()
	}

	if !git.ExistOnOrigin(currentBranch) {
		println("Branch " + color.HiGreenString(currentBranch) + " does not exist on origin")
		println("Please push your branch before release")

		println("Run " + color.HiGreenString("'cominnek publish'") + " to push your branch")
		os.Exit(1)
	}
}
