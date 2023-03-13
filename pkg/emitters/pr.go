package emitters

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	eventNames "github.com/Minnek-Digital-Studio/cominnek/pkg/events/names"
)

var pullRequestNames = new(eventNames.PullRequest)

type PullRequest struct {
}

func (c *PullRequest) Init() {
	events.App.Emit(pullRequestNames.Init())
}

func (c *PullRequest) Failed(error string) {
	events.App.Emit(pullRequestNames.Failed(), error)
}

func (c *PullRequest) Success(message string) {
	events.App.Emit(pullRequestNames.Success(), message)
}
