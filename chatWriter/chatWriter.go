package chatwriter

import "github.com/vivanshah/hangout-json-parser/models"

// ChatWriter interface to support multiple output types
type ChatWriter interface {
	WriteChat(chat models.Chat) error
}
