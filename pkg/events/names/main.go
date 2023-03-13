package eventNames

import "github.com/kataras/go-events"

func builder(base string, event string) events.EventName {
	return events.EventName(base + ":" + event)
}