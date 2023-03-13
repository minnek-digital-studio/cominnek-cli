package events

import eventNames "github.com/Minnek-Digital-Studio/cominnek/pkg/events/names"

var commitNames = new(eventNames.Commit)
var pushNames = new(eventNames.Push)

func Watcher(){
	App.On("init:root", func(payload ...interface{}) {
		// data := payload[0].(*emitterTypes.IRootEmitter)
	}); 

	App.On(commitNames.Init(), func(payload ...interface{}) {
		println(commitNames.Init())
	});
	App.On(commitNames.Failed(), func(payload ...interface{}) {
		err := payload[0].(string)

		println(commitNames.Failed())
		println(err)
	});
	App.On(commitNames.Success(), func(payload ...interface{}) {
		message := payload[0].(string)

		println(commitNames.Success())
		println(message)
	});

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
}