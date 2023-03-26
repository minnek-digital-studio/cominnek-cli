package emitters

import (
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	eventNames "github.com/Minnek-Digital-Studio/cominnek/pkg/events/names"
)

var resetNames = new(eventNames.Reset)

type Reset struct {
}

func (c *Reset) Init(data string) {
	events.App.Emit(resetNames.Init(), data)
}

func (c *Reset) Failed(error string) {
	events.App.Emit(resetNames.Failed(), error)
}

func (c *Reset) Success(data string) {
	events.App.Emit(resetNames.Success(), data)
}
