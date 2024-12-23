package lava

type Lava struct {
	Version string
}

type Module interface {
	OnLoad()
	OnEnable()
	OnDisable()
	Name() string
}

var LavaAPI Lava = Lava{"1.0"}
