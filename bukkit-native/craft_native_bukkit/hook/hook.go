package hook

import (
	"github.com/iRedTea/lava/bukkit-native/craft_native_bukkit"
	"unsafe"
)
import "C"

var HOOKS BukkitHooks

//export LoadNativeBukkit
func LoadNativeBukkit(bukkitVersion string, hooks BukkitHooks) unsafe.Pointer {
	HOOKS = hooks
	info := craft_native_bukkit.NewCraftBukkitInfo(bukkitVersion)
	return unsafe.Pointer(craft_native_bukkit.NewCraftBukkit(info))
}
