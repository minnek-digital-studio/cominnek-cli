package pkg_action

import (
	"errors"
	"strconv"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Minnek-Digital-Studio/cominnek/config"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/ask"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
)

var haveNumber = func() bool { return config.AppData.Reset.Number != "" }
var haveType = func() bool { return config.AppData.Reset.Type != "" }
var haveCommit = func() bool { return config.AppData.Reset.Commit != "" }

var resetEmitter = new(emitters.Reset)

func showConfirmation() {
	if config.AppData.Reset.Confirm {
		return
	}

	if haveNumber() && haveType() {
		ask.One(&survey.Confirm{
			Message: "Are you sure you want to reset " + config.AppData.Reset.Type + " " + config.AppData.Reset.Number + " commits?",
		}, &config.AppData.Reset.Confirm, survey.WithValidator(survey.Required))
	}

	if haveCommit() && haveType() {
		ask.One(&survey.Confirm{
			Message: "Are you sure you want to reset " + config.AppData.Reset.Type + " to commit " + config.AppData.Reset.Commit + "?",
		}, &config.AppData.Reset.Confirm, survey.WithValidator(survey.Required))
	}

	if config.AppData.Reset.Type == "merge" {
		ask.One(&survey.Confirm{
			Message: "Are you sure you want to reset merge?",
		}, &config.AppData.Reset.Confirm, survey.WithValidator(survey.Required))
	}
}

func resetQuestions() {

	git.GetAllCommits()

	if (haveNumber() || haveCommit()) && !haveType() {
		config.AppData.Reset.Type = "soft"
	}

	if !haveType() {
		ask.One(&survey.Select{
			Message: "Select a reset type:",
			Options: []string{"soft", "mixed", "hard", "merge", "keep"},
		}, &config.AppData.Reset.Type, survey.WithValidator(survey.Required))
	}

	if config.AppData.Reset.Type == "merge" {
		showConfirmation()
		return
	}

	if !haveNumber() && !haveCommit() {
		ask.One(&survey.Select{
			Message: "Select a reset target:",
			Options: []string{"commit", "number"},
		}, &config.AppData.Reset.Target, survey.WithValidator(survey.Required))
	}

	if config.AppData.Reset.Target == "number" && !haveNumber() {
		ask.One(&survey.Input{
			Message: "Enter the number of commits to reset:",
		}, &config.AppData.Reset.Number, survey.WithValidator(survey.Required),
			survey.WithValidator(func(val interface{}) error {
				if _, err := strconv.Atoi(val.(string)); err != nil {
					resetEmitter.Failed("please enter a number")
					return errors.New("please enter a number")
				}
				return nil
			},
			))
	}

	if config.AppData.Reset.Target == "commit" && !haveCommit() {
		ask.One(&survey.Select{
			Message: "Select a commit to reset to:",
			Options: git.GetAllCommits(),
		}, &config.AppData.Reset.Commit, survey.WithValidator(survey.Required))

		config.AppData.Reset.Commit = git.GetCommitHash(config.AppData.Reset.Commit)
	}

	showConfirmation()
}

func Reset() {
	resetQuestions()
	resetEmitter.Init("Reset type: " + config.AppData.Reset.Type)

	if config.AppData.Reset.Confirm {
		r_type := config.AppData.Reset.Type
		r_number := config.AppData.Reset.Number
		r_commit := config.AppData.Reset.Commit

		git.Reset(r_type, r_number, r_commit)
		since := time.Since(config.AppData.Start).String()

		resetEmitter.Success("Reset successful (" + since + ")")
		println("Reset successful (" + since + ")")
	}
}
