package hook

import (
	"github.com/google/uuid"
	"github.com/iRedTea/lava/bukkit-native/craft_native_bukkit"
	"unsafe"
)
import "C"

//export LoadNativeBukkit
func LoadNativeBukkit(bukkitVersion string, hooksPointer unsafe.Pointer) unsafe.Pointer {
	SendMessageHookPtr = (*func(playerId uuid.UUID, message string))(hooksPointer)
	info := craft_native_bukkit.NewCraftBukkitInfo(bukkitVersion)
	return unsafe.Pointer(craft_native_bukkit.NewCraftBukkit(info))
}
