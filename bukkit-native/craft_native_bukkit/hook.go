package craft_native_bukkit

import (
	"github.com/google/uuid"
	"unsafe"
)
import "C"

type BukkitHooks interface {
	SendMessageToPlayerHook(playerId uuid.UUID, message string)
}

var HOOKS BukkitHooks

//export LoadNativeBukkit
func LoadNativeBukkit(bukkitVersion string, hooks BukkitHooks) unsafe.Pointer {
	HOOKS = hooks
	info := NewCraftBukkitInfo(bukkitVersion)
	return unsafe.Pointer(NewCraftBukkit(info))
}
