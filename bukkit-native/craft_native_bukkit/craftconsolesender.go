package craft_native_bukkit

import (
	"fmt"
	bukkit_native "github.com/iRedTea/lava/bukkit-native"
)

type CraftConsoleSender struct {
	bukkit_native.CommandSender
}

func (c *CraftConsoleSender) Name() string {
	return "CONSOLE"
}

func (c *CraftConsoleSender) SendMessage(message string) {
	fmt.Println(message)
}
