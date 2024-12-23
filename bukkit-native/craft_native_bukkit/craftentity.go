package craft_native_bukkit

import (
	"github.com/google/uuid"
	bukkit "github.com/iRedTea/lava/bukkit-native"
)

type CraftEntity struct {
	bukkit.Entity
	uniqueId uuid.UUID
	name     string
}
