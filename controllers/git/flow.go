package git_controller

import (
	"fmt"

	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

type FlowResponse interface {
	Commit() string
	Message() string
	IgnoreOutput() bool
}

type flowResponse struct {
	command string
	message string
	ignoreOutput bool
}

func Feature(ticket string) string {
	return "git flow feature start " + ticket
}

func Hotfix(ticket string) []flowResponse {
	return getCommandsFlow("git flow hotfix start " + ticket)
}

func Release(ticket string) []flowResponse {
	return getCommandsFlow("git flow release start " + ticket)
}

func Support(ticket string) []flowResponse {
	return getCommandsFlow("git flow support start " + ticket)
}

func Init() string {
	return "git flow init"
}

func getCommandsFlow(mainCmd string) []flowResponse {
	flowRes := []flowResponse{}
	branch := "develop"

	if !CheckIfBranch(branch) {
		fmt.Println("Switching to develop branch")
		shell.ExecuteCommand("git checkout " + branch)
		flowRes = append(flowRes, flowResponse{
			command: "git checkout " + branch,
			message: "switching to " + branch,
			ignoreOutput: true,
		})
	}

	shell.ExecuteCommand("git fetch origin");
	flowRes = append(flowRes, flowResponse{
		command: "git fetch origin",
		message: "fetching data from origin",
	})

	if !CheckChangesFromOrigin() {
		shell.ExecuteCommand(Pull())
		flowRes = append(flowRes, flowResponse{
			command: "git pull origin " + branch,
			message: "pulling data...",
		})
	}

	shell.ExecuteCommand(mainCmd)
	flowRes = append(flowRes, flowResponse{
		command: mainCmd,
		message: "",
	})

	return flowRes
}

func (fR *flowResponse) Command() string {
	return fR.command
}

func (fR *flowResponse) Message() string {
	return fR.message
}

func (fR *flowResponse) IgnoreOutput() bool {
	return fR.ignoreOutput
}