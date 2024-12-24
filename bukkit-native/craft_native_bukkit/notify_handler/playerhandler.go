package notify_handler

import (
	"github.com/google/uuid"
	"github.com/iRedTea/lava/bukkit-native/craft_native_bukkit"
)

func PlayerJoinHandler(uniqueId uuid.UUID, name string, message string) string {
	return craft_native_bukkit.Instance.Players().(*craft_native_bukkit.CraftPlayerList).Join(uniqueId, name, message)
}

func PlayerQuitHandler(uniqueId uuid.UUID, message string) string {
	return craft_native_bukkit.Instance.Players().(*craft_native_bukkit.CraftPlayerList).Quit(uniqueId, message)
}
