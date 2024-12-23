package bukkit_native

import (
	"github.com/google/uuid"
)

type OfflinePlayer interface {
	Entity
	IsOnline() bool
}

type Player interface {
	OfflinePlayer
	CommandSender
}

type PlayerList interface {
	All() []Player
	Get(id uuid.UUID) OfflinePlayer
	GetByName(name string) OfflinePlayer
	AllOffline() []OfflinePlayer
}
