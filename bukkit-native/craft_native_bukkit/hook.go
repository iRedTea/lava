package craft_native_bukkit

import (
	"fmt"
	"github.com/google/uuid"
	bukkit "github.com/iRedTea/lava/bukkit-native"
)

type BukkitHooks interface {
	SendMessageToPlayerHook(playerId uuid.UUID, message string)
}

var HOOKS BukkitHooks

func LoadNativeBukkit(bukkitVersion string, hooks BukkitHooks) bukkit.Bukkit {
	fmt.Println("Loading Native Bukkit from lava... Version of bukkit: " + bukkitVersion)
	HOOKS = hooks
	if HOOKS == nil {
		fmt.Println("hooks could not be nil")
		return nil
	}
	bukkitInfo := NewCraftBukkitInfo(bukkitVersion)
	if bukkitInfo == nil {
		fmt.Println("bukkitInfo could not be nil")
		return nil
	}
	return NewCraftBukkit(bukkitInfo)
}
