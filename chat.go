package main

import "time"

type Chat struct {
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
}
