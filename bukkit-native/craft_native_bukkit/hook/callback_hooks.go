package hook

import "github.com/google/uuid"

type BukkitHooks interface {
	SendMessageToPlayerHook(playerId uuid.UUID, message string)
}
