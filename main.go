package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var input = flag.String("input", "hangouts.json", "Input json file")
	flag.Parse()

	jsonFile, err := os.Open(*input)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var hangouts Hangouts
	fmt.Println("Parsing hangouts file")
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(byteValue, &hangouts)

	fmt.Println("Loaded ", len(hangouts.Conversations), " conversations")

	chats := []Chat{}
	for _, c := range hangouts.Conversations {
		chat := Chat{
			ParticipantIDs:   map[string]int{},
			ParticipantNames: map[string]string{},
			Messages:         []Message{},
		}

		participants := c.Conversation.Conversation.ParticipantData
		for n, p := range participants {
			chat.ParticipantIDs[p.ID.ChatID] = n
			chat.ParticipantNames[p.ID.ChatID] = p.FallbackName
			chat.Title = chat.Title + p.FallbackName + ", "
		}
		chat.Title = strings.TrimRight(chat.Title, ", ")
		for _, e := range c.Events {
			message := Message{}
			t, _ := strconv.ParseInt(e.Timestamp, 10, 64)
			message.Timestamp = time.Unix(0, t)
			message.SenderID = chat.ParticipantIDs[e.SenderID.ChatID]
			message.Sender = chat.ParticipantNames[e.SenderID.ChatID]
			for _, s := range e.ChatMessage.MessageContent.Segment {
				if s.Type == "TEXT" {
					message.Text = message.Text + " " + s.Text
				}
			}
		}
		chats = append(chats, chat)
	}

	for _, chat := range chats {
		fmt.Println(chat.Title)
	}
}
