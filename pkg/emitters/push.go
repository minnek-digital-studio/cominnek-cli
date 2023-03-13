package emitters

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	eventNames "github.com/Minnek-Digital-Studio/cominnek/pkg/events/names"
)

var pushNames = new(eventNames.Push)

type Push struct {
}

func (c *Push) Init() {
	events.App.Emit(pushNames.Init())
}

func (c *Push) Failed(error string) {
	events.App.Emit(pushNames.Failed(), error)
}

func (c *Push) Success(message string) {
	events.App.Emit(pushNames.Success(), message)
}
