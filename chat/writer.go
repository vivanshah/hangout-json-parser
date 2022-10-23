package chat

// Writer interface to support multiple output types
type Writer interface {
	WriteChat(conversation Conversation) error
}
