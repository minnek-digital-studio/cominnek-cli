package eventNames

import "github.com/kataras/go-events"

func _buildEventName(event string) events.EventName {
	return builder("commit", event)
}

type Commit struct {
}

func (c *Commit) Init()  events.EventName {
	return _buildEventName("init")
}

func (c *Commit) Failed() events.EventName{
	return _buildEventName("failed")
}

func (c *Commit) Success()  events.EventName {
	return _buildEventName("success")
}
