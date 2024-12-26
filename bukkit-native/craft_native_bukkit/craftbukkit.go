package craft_native_bukkit

import (
	bukkit "github.com/iRedTea/lava/bukkit-native"
	"github.com/iRedTea/lava/bukkit-native/events"
)

type CraftBukkit struct {
	bukkit.Bukkit
	bukkitInfo    bukkit.BukkitInfo
	registry      events.Registry
	players       bukkit.PlayerList
	commandSender bukkit.CommandSender
}

func (c *CraftBukkit) BukkitInfo() bukkit.BukkitInfo {
	return c.bukkitInfo
}
func (c *CraftBukkit) Players() bukkit.PlayerList {
	return c.players
}
func (c *CraftBukkit) ConsoleCommandSender() bukkit.CommandSender {
	return c.commandSender
}

var Instance *CraftBukkit

func NewCraftBukkit(newBukkitInfo *CraftBukkitInfo) *CraftBukkit {
	Instance = &CraftBukkit{
		bukkitInfo:    newBukkitInfo,
		registry:      &CraftRegistry{},
		players:       NewCraftPlayerList(),
		commandSender: &CraftConsoleSender{},
	}
	return Instance
}

type CraftBukkitInfo struct {
	bukkit.BukkitInfo
	version string
}

func (c *CraftBukkitInfo) BukkitVersion() string {
	return c.version
}

func NewCraftBukkitInfo(newVersion string) CraftBukkitInfo {
	return CraftBukkitInfo{version: newVersion}
}
