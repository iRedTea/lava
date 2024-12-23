package craft_native_bukkit

import (
	"github.com/iRedTea/lava/bukkit-native/events"
	"github.com/iRedTea/lava/bukkit-native/events/event_types"
)

type CraftRegistry struct {
	events.Registry
	handlers map[event_types.EventType][]func(event *events.Event)
}

func NewCraftRegistry() *CraftRegistry {
	return &CraftRegistry{
		handlers: make(map[event_types.EventType][]func(event *events.Event)),
	}
}

func (r *CraftRegistry) Register(handler func(event *events.Event), eventType event_types.EventType) {
	r.handlers[eventType] = append(r.handlers[eventType], handler)
}

func (r *CraftRegistry) Invoke(event *events.Event) {
	handlers, exists := r.handlers[event.Type]
	if !exists {
		return
	}

	for _, handler := range handlers {
		handler(event)
	}
}
