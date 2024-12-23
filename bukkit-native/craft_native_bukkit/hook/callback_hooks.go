package hook

import "C"
import "github.com/google/uuid"

var SendMessageHookPtr *func(playerId uuid.UUID, message string)

func SendMessageToPlayerHook(playerId uuid.UUID, message string) {
	(*SendMessageHookPtr)(playerId, message)
}
