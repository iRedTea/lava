package bukkit_native

type Bukkit interface {
	Players() PlayerList
	BukkitInfo() BukkitInfo
	ConsoleCommandSender() CommandSender
}

type BukkitInfo interface {
	BukkitVersion() string
}
