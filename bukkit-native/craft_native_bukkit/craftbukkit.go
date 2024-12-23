package craft_native_bukkit

import (
	bukkit "github.com/iRedTea/lava/bukkit-native"
	"github.com/iRedTea/lava/bukkit-native/events"
)

var consoleCommandSender = &CraftConsoleSender{}

type CraftBukkit struct {
	bukkit.Bukkit
	bukkitInfo bukkit.BukkitInfo
	registry   events.Registry
}

func (c *CraftBukkit) BukkitInfo() bukkit.BukkitInfo {
	return c.bukkitInfo
}

var Instance *CraftBukkit

func NewCraftBukkit(bukkitInfo *CraftBukkitInfo) *CraftBukkit {
	Instance = &CraftBukkit{
		bukkitInfo: bukkitInfo,
		registry:   &CraftRegistry{},
	}
	return Instance
}

type CraftBukkitInfo struct {
	bukkit.BukkitInfo
	version string
}

func (c *CraftBukkitInfo) GetVersion() string {
	return c.version
}

func (c *CraftBukkitInfo) ConsoleCommandSender() bukkit.CommandSender {
	return consoleCommandSender
}

func NewCraftBukkitInfo(version string) *CraftBukkitInfo {
	return &CraftBukkitInfo{version: version}
}

func initialize() {

}
