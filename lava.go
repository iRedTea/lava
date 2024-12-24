package lava

import bukkit "github.com/iRedTea/lava/bukkit-native"

type Lava struct {
	Version string
	Bukkit  bukkit.Bukkit
}

type Module interface {
	OnLoad()
	OnEnable()
	OnDisable()
	Name() string
}

var LavaAPI *Lava = &Lava{}
