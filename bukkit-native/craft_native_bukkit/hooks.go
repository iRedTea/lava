package craft_native_bukkit

import (
	"github.com/google/uuid"
	"unsafe"
)
import "C"

//export SendMessageToPlayerHook
func SendMessageToPlayerHook(playerId uuid.UUID, message string) {
	// call java code
}

//export LoadNativeBukkit
func LoadNativeBukkit(bukkitVersion string) unsafe.Pointer {
	info := NewCraftBukkitInfo(bukkitVersion)
	return unsafe.Pointer(NewCraftBukkit(info))
}
