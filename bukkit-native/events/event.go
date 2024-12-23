package events

import "github.com/iRedTea/lava/bukkit-native/events/event_types"

type Event struct {
	Type event_types.EventType
}

type Cancellable interface {
	Cancel()
	IsCancelled() bool
}

type Registry interface {
	Register(handler func(event *Event), eventType event_types.EventType)
	Invoke(event *Event)
}
