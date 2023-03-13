package eventNames

import "github.com/kataras/go-events"

func _buildPushtName(event string) events.EventName {
	return builder("push", event)
}

type Push struct {
}

func (c *Push) Init()  events.EventName {
	return _buildPushtName("init")
}

func (c *Push) Failed() events.EventName{
	return _buildPushtName("failed")
}

func (c *Push) Success()  events.EventName {
	return _buildPushtName("success")
}
