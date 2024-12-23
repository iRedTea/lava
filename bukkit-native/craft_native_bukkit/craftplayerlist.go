package craft_native_bukkit

import (
	"github.com/google/uuid"
	bukkit "github.com/iRedTea/lava/bukkit-native"
	"github.com/iRedTea/lava/bukkit-native/events"
	"slices"
)

type CraftPlayerList struct {
	bukkit.PlayerList
	onlinePlayersById map[uuid.UUID]*CraftPlayer
	nameToUUID        map[string]uuid.UUID
	playersCache      []bukkit.Player
}

func NewCraftPlayerList() *CraftPlayerList {
	return &CraftPlayerList{}
}

func (c *CraftPlayerList) All() []bukkit.Player {
	return c.playersCache
}

// return join message
func (c *CraftPlayerList) join(uniqueId uuid.UUID, name string, message string) string {
	player := NewCraftPlayer(uniqueId, name)
	player.IsOnline_ = true
	c.onlinePlayersById[uniqueId] = player
	c.nameToUUID[name] = uniqueId

	c.playersCache = append(c.playersCache, player)

	joinEvent := events.PlayerJoinEvent{
		PlayerEvent: events.PlayerEvent{},
		JoinMessage: message,
	}
	var event = joinEvent.Event
	Instance.registry.Invoke(&event)
	return joinEvent.JoinMessage
}

// return quit message
func (c *CraftPlayerList) quit(uniqueId uuid.UUID, message string) string {
	c.onlinePlayersById[uniqueId].IsOnline_ = false
	delete(c.onlinePlayersById, uniqueId)
	i := 0
	for j, p := range c.playersCache {
		if p.UniqueID() == uniqueId {
			i = j
			break
		}
	}
	c.playersCache = slices.Delete(c.playersCache, i, i+1)

	quitEvent := events.PlayerQuitEvent{
		PlayerEvent: events.PlayerEvent{},
		QuitMessage: message,
	}
	var event = quitEvent.Event
	Instance.registry.Invoke(&event)

	return quitEvent.QuitMessage
}
