package bukkit_native

type Bukkit interface {
	Players() PlayerList
	BukkitInfo() BukkitInfo
}

type BukkitInfo interface {
	BukkitVersion() string
}
