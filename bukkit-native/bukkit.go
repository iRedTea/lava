package bukkit_native

type Bukkit interface {
	Players() PlayerList
}

type BukkitInfo interface {
	BukkitVersion() string
}
