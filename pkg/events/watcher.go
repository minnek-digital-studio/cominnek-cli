package events

func Watcher(){
	App.On("init:root", func(payload ...interface{}) {
		// data := payload[0].(*emitterTypes.IRootEmitter)
	}); 
}