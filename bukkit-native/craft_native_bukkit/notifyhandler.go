package craft_native_bukkit

import "github.com/google/uuid"

func playerJoinHandler(uniqueId uuid.UUID, name string, message string) string {
	return Instance.players.(*CraftPlayerList).join(uniqueId, name, message)
}

func playerQuitHandler(uniqueId uuid.UUID, message string) string {
	return Instance.players.(*CraftPlayerList).quit(uniqueId, message)
}
