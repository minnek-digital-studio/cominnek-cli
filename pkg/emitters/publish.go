package emitters

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	eventNames "github.com/Minnek-Digital-Studio/cominnek/pkg/events/names"
)

var publishNames = new(eventNames.Publish)

type Publish struct {
}

func (c *Publish) Init() {
	events.App.Emit(publishNames.Init())
}

func (c *Publish) Failed(error string) {
	events.App.Emit(publishNames.Failed(), error)
}

func (c *Publish) Success(message string) {
	events.App.Emit(publishNames.Success(), message)
}
