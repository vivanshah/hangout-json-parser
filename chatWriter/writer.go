package chat

import "github.com/vivanshah/hangout-json-parser/models"

// Writer interface to support multiple output types
type Writer interface {
	WriteChat(chat models.Chat) error
}
