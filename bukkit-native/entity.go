package bukkit_native

import "github.com/google/uuid"

type Entity interface {
	UniqueID() uuid.UUID
}
