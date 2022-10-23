package chat

import "time"

type Conversation struct {
	Title            string
	ParticipantIDs   map[string]int
	ParticipantNames map[string]string
	Messages         []Message
}

type Message struct {
	Sender    string
	SenderID  int
	Timestamp time.Time
	Text      string
	Self      bool
	ImageURLs []string
}

func (c Conversation) String() string {
	return c.Title
}
