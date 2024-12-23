package events

import (
	b "github.com/iRedTea/lava/bukkit-native"
	"github.com/iRedTea/lava/bukkit-native/events/event_types"
)

type PlayerEvent struct {
	Player b.Player
}

// JOIN

type PlayerJoinEvent struct {
	JoinMessage string
}

func (e *PlayerJoinEvent) Type() event_types.EventType {
	return event_types.PlayerJoinEvent
}

// QUIT

type PlayerQuitEvent struct {
	QuitMessage string
}

func (e *PlayerQuitEvent) Type() event_types.EventType {
	return event_types.PlayerQuitEvent
}
