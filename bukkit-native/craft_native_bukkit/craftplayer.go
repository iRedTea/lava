package craft_native_bukkit

import (
	"github.com/google/uuid"
	bukkit "github.com/iRedTea/lava/bukkit-native"
)

type CraftPlayer struct {
	CraftEntity
	bukkit.Player
	IsOnline_ bool
}

func (c *CraftPlayer) IsOnline() bool {
	return c.IsOnline_
}

func (c *CraftPlayer) SendMessage(message string) {
	go HOOKS.SendMessageToPlayerHook(c.UniqueID(), message)
}

func NewCraftPlayer(uniqueId uuid.UUID, name string) *CraftPlayer {
	return &CraftPlayer{
		IsOnline_: false,
		CraftEntity: CraftEntity{
			uniqueId: uniqueId,
			name:     name,
		}}
}
