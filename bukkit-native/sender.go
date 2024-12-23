package bukkit_native

type CommandSender interface {
	Name() string
	SendMessage(message string)
}
