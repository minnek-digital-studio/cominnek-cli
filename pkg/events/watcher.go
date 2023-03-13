package events

import (
	"fmt"

	emitterTypes "github.com/Minnek-Digital-Studio/cominnek/pkg/emitters/types"
	eventNames "github.com/Minnek-Digital-Studio/cominnek/pkg/events/names"
)

var commitNames = new(eventNames.Commit)
var pushNames = new(eventNames.Push)
var branchNames = new(eventNames.Branch)
var publishNames = new(eventNames.Publish)

func Watcher() {
	App.On("init:root", func(payload ...interface{}) {
		// data := payload[0].(*emitterTypes.IRootEmitter)
	})

	//! Commit
	App.On(commitNames.Init(), func(payload ...interface{}) {
		println(commitNames.Init())
	})
	App.On(commitNames.Failed(), func(payload ...interface{}) {
		err := payload[0].(string)

		println(commitNames.Failed())
		println(err)
	})
	App.On(commitNames.Success(), func(payload ...interface{}) {
		message := payload[0].(string)

		println(commitNames.Success())
		println(message)
	})

	//! Push
	App.On(pushNames.Init(), func(payload ...interface{}) {
		println(pushNames.Init())
	})
	App.On(pushNames.Failed(), func(payload ...interface{}) {
		err := payload[0].(string)

		println(pushNames.Failed())
		println(err)
	})
	App.On(pushNames.Success(), func(payload ...interface{}) {
		message := payload[0].(string)

		println(pushNames.Success())
		println(message)
	})

	//! Branch
	App.On(branchNames.Init(), func(payload ...interface{}) {
		data := payload[0].(emitterTypes.IBranchEventData)
		println(branchNames.Init())
		fmt.Printf("%+v\n", data)
	})
	App.On(branchNames.Failed(), func(payload ...interface{}) {
		data := payload[0].(emitterTypes.IBranchFailedData)
		
		println(branchNames.Failed())
		fmt.Printf("%+v\n", data)
	})
	App.On(branchNames.Success(), func(payload ...interface{}) {
		data := payload[0].(emitterTypes.IBranchEventData)

		println(branchNames.Success())
		fmt.Printf("%+v\n", data)
	})

	//! Publish
	App.On(publishNames.Init(), func(payload ...interface{}) {
		println(publishNames.Init())
	})
	App.On(publishNames.Failed(), func(payload ...interface{}) {
		err := payload[0].(string)

		println(publishNames.Failed())
		println(err)
	})
	App.On(publishNames.Success(), func(payload ...interface{}) {
		message := payload[0].(string)

		println(publishNames.Success())
		println(message)
	})
}
