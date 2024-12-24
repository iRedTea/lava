package lava

import bukkit "github.com/iRedTea/lava/bukkit-native"

type Lava struct {
	Version     string
	Bukkit      bukkit.Bukkit
	TempContext Context
}

type Context struct {
	OnEnable  func()
	OnDisable func()
	OnLoad    func()
	Name      func() string
}

type Module interface {
	OnLoad()
	OnEnable()
	OnDisable()
	Name() string
}

var LavaAPI *Lava = &Lava{}
