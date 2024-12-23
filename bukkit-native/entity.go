package bukkit_native

import "github.com/google/uuid"

type Entity interface {
	Name() string
	UniqueID() uuid.UUID
}
