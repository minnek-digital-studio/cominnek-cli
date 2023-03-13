package eventNames

import "github.com/kataras/go-events"

func _buildPullRequestName(event string) events.EventName {
	return builder("pr", event)
}

type PullRequest struct {
}

func (c *PullRequest) Init()  events.EventName {
	return _buildPullRequestName("init")
}

func (c *PullRequest) Failed() events.EventName{
	return _buildPullRequestName("failed")
}

func (c *PullRequest) Success()  events.EventName {
	return _buildPullRequestName("success")
}
