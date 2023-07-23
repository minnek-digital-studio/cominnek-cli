package github

import (
	"log"

	"github.com/Minnek-Digital-Studio/cominnek/config"
	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/loading"
	"github.com/Minnek-Digital-Studio/cominnek/controllers/project"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/emitters"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/git"
)

var publishEmitter = new(emitters.Publish)

func _checkBranch() []string {
	var branch []string
	currentBranch := git_controller.GetCurrentBranch()
	branchType := git_controller.GetBranchType()
	branchData := project.GetConfigByName(branchType)
	config.AppData.Branch.Data = branchData

	if currentBranch == "master" {
		log.Fatal("You can't create a pull request from the master branch")
	}

	if len(branchData.To) > 0 {
		branch = append(branch, branchData.To...)
	} else {
		branch = append(branch, branchData.From)
	}

	return branch
}

func NewCreatePullRequest(ticket string, baseBranch string) {
	loading.Start("Checking branch ")
	branches := _checkBranch()
	loading.Stop()

	for _, branch := range branches {
		CreatePullRequest(ticket, branch)
	}
}

func Publish(ticket string) {
	git.PushPublish()
	NewCreatePullRequest(ticket, "")
	publishEmitter.Success("Publish complete")
	log.Println("Publish complete")
}
