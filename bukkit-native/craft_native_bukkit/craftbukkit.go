package craft_native_bukkit

import (
	"fmt"
	bukkit "github.com/iRedTea/lava/bukkit-native"
	"github.com/iRedTea/lava/bukkit-native/events"
)

var ConsoleCommandSender bukkit.CommandSender = &CraftConsoleSender{}

type CraftBukkit struct {
	bukkit.Bukkit
	bukkitInfo  bukkit.BukkitInfo
	registry    events.Registry
	players     bukkit.PlayerList
	fingerprint string
}

func (c *CraftBukkit) BukkitInfo() bukkit.BukkitInfo {
	return c.bukkitInfo
}
func (c *CraftBukkit) Players() bukkit.PlayerList {
	return c.players
}
func (c *CraftBukkit) ConsoleCommandSender() bukkit.CommandSender {
	fmt.Println("[1] SENDER CHECK 1")
	fmt.Println("[1] SENDER NAME " + ConsoleCommandSender.Name())
	return ConsoleCommandSender
}

func (c *CraftBukkit) FPR() string {
	return c.fingerprint
}

var Instance *CraftBukkit

func NewCraftBukkit(newBukkitInfo *CraftBukkitInfo) *CraftBukkit {
	Instance = &CraftBukkit{
		bukkitInfo:  newBukkitInfo,
		registry:    &CraftRegistry{},
		players:     NewCraftPlayerList(),
		fingerprint: "callback231",
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
