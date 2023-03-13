package eventNames

import "github.com/kataras/go-events"

func _buildPublishName(event string) events.EventName {
	return builder("publish", event)
}

type Publish struct {
}

func (c *Publish) Init()  events.EventName {
	return _buildPublishName("init")
}

func (c *Publish) Failed() events.EventName{
	return _buildPublishName("failed")
}

func (c *Publish) Success()  events.EventName {
	return _buildPublishName("success")
}
