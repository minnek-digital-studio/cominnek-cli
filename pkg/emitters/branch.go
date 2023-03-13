package emitters

import (
	emitterTypes "github.com/Minnek-Digital-Studio/cominnek/pkg/emitters/types"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/events"
	eventNames "github.com/Minnek-Digital-Studio/cominnek/pkg/events/names"
)

var branchNames = new(eventNames.Branch)

type Branch struct {
}

func (c *Branch) Init(data emitterTypes.IBranchEventData) {
	events.App.Emit(branchNames.Init(), data)
}

func (c *Branch) Failed(error emitterTypes.IBranchFailedData) {
	events.App.Emit(branchNames.Failed(), error)
}

func (c *Branch) Success(data emitterTypes.IBranchEventData) {
	events.App.Emit(branchNames.Success(), data)
}
