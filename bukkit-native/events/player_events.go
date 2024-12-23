package events

import (
	bukkit "github.com/iRedTea/lava/bukkit-native"
)

type PlayerEvent struct {
	Event
	Player bukkit.Player
}

type PlayerJoinEvent struct {
	PlayerEvent
	JoinMessage string
}

type PlayerQuitEvent struct {
	PlayerEvent
	QuitMessage string
}
