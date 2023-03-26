package emitters

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	eventNames "github.com/Minnek-Digital-Studio/cominnek/pkg/events/names"
)

var commitNames = new(eventNames.Commit)

type Commit struct {
}

func (c *Commit) Init() {
	events.App.Emit(commitNames.Init())
}

func (c *Commit) Failed(error string) {
	events.App.Emit(commitNames.Failed(), error)
}

func (c *Commit) Success(message string) {
	events.App.Emit(commitNames.Success(), message)
}
