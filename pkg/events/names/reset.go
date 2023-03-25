package eventNames

import "github.com/kataras/go-events"

func _buildResetName(event string) events.EventName {
	return builder("reset", event)
}

type Reset struct {
}

func (c *Reset) Init()  events.EventName {
	return _buildResetName("init")
}

func (c *Reset) Failed() events.EventName{
	return _buildResetName("failed")
}

func (c *Reset) Success()  events.EventName {
	return _buildResetName("success")
}
