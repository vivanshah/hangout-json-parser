package chatWriter
import "github.com/vivanshah/hangout-json-parser/models"

type ChatWriter interface {
	WriteChat(chat models.Chat) error
}