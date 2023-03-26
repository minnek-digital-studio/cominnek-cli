package events

import (
	eventNames "github.com/Minnek-Digital-Studio/cominnek/pkg/events/names"
)

var commitNames = new(eventNames.Commit)
var pushNames = new(eventNames.Push)
var branchNames = new(eventNames.Branch)
var publishNames = new(eventNames.Publish)
var pullRequest = new(eventNames.PullRequest)
var resetNames = new(eventNames.Reset)

func Watcher() {
	App.On("init:root", func(payload ...interface{}) {})

	//! Commit
	App.On(commitNames.Init(), func(payload ...interface{}) {})
	App.On(commitNames.Failed(), func(payload ...interface{}) {})
	App.On(commitNames.Success(), func(payload ...interface{}) {})

	//! Push
	App.On(pushNames.Init(), func(payload ...interface{}) {})
	App.On(pushNames.Failed(), func(payload ...interface{}) {})
	App.On(pushNames.Success(), func(payload ...interface{}) {})

	//! Branch
	App.On(branchNames.Init(), func(payload ...interface{}) {})
	App.On(branchNames.Failed(), func(payload ...interface{}) {})
	App.On(branchNames.Success(), func(payload ...interface{}) {})

	//! Publish
	App.On(publishNames.Init(), func(payload ...interface{}) {})
	App.On(publishNames.Failed(), func(payload ...interface{}) {})
	App.On(publishNames.Success(), func(payload ...interface{}) {})

	//! Pull Request
	App.On(pullRequest.Init(), func(payload ...interface{}) {})
	App.On(pullRequest.Failed(), func(payload ...interface{}) {})
	App.On(pullRequest.Success(), func(payload ...interface{}) {})

	//! Reset
	App.On(resetNames.Init(), func(payload ...interface{}) {})
	App.On(resetNames.Failed(), func(payload ...interface{}) {})
	App.On(resetNames.Success(), func(payload ...interface{}) {})
}
