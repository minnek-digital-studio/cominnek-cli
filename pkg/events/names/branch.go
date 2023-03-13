package eventNames

import "github.com/kataras/go-events"

func _buildBranchName(event string) events.EventName {
	return builder("branch", event)
}

type Branch struct {
}

func (c *Branch) Init()  events.EventName {
	return _buildBranchName("init")
}

func (c *Branch) Failed() events.EventName{
	return _buildBranchName("failed")
}

func (c *Branch) Success()  events.EventName {
	return _buildBranchName("success")
}
