package bukkit_native

type CommandSender interface {
	SendMessage(message string) string
	SendRawMessage(raw string) string
}
