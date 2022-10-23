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

	"github.com/kennygrant/sanitize"
	"github.com/vivanshah/hangout-json-parser/chat"
	"github.com/vivanshah/hangout-json-parser/models"

	"github.com/AlecAivazis/survey/v2"
)

func main() {

	var (
		input = flag.String("input", "hangouts.json", "Input json file")
		//outputFormat = flag.String("f","txt", "Output file format")
	)
	flag.Parse()

	jsonFile, err := os.Open(*input)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	var hangouts models.Hangouts
	fmt.Println("Parsing hangouts file")
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteValue, &hangouts)
	if err != nil {
		fmt.Println("Error parsing input file: ", err.Error())
		os.Exit((1))
	}

	fmt.Println("Loaded ", len(hangouts.Conversations), " conversations")
	chatMap := map[string]chat.Conversation{}
	chatTitles := []string{}
	chats := []chat.Conversation{}
	for _, c := range hangouts.Conversations {
		convo := chat.Conversation{
			ParticipantIDs:   map[string]int{},
			ParticipantNames: map[string]string{},
			Messages:         []chat.Message{},
		}

		participants := c.Conversation.Conversation.ParticipantData
		for n, p := range participants {
			convo.ParticipantIDs[p.ID.ChatID] = n
			convo.ParticipantNames[p.ID.ChatID] = p.FallbackName
			convo.Title = convo.Title + p.FallbackName + ", "
		}
		convo.Title = strings.TrimRight(convo.Title, ", ")
		for _, e := range c.Events {
			if e.EventType != "REGULAR_CHAT_MESSAGE" {
				continue
			}
			message := chat.Message{}
			t, _ := strconv.ParseInt(e.Timestamp, 10, 64)
			t = t * 1000
			message.Timestamp = time.Unix(0, t)
			message.SenderID = convo.ParticipantIDs[e.SenderID.ChatID]
			message.Sender = convo.ParticipantNames[e.SenderID.ChatID]
			message.Self = e.SenderID.ChatID == e.SelfEventState.UserID.ChatID
			for _, s := range e.ChatMessage.MessageContent.Segment {
				message.Text = message.Text + " " + s.Text
				if s.Type == "LINK" {
					message.Text = message.Text + "[" + s.LinkData.LinkTarget + "]"
				}
			}
			for _, a := range e.ChatMessage.MessageContent.Attachment {
				message.ImageURLs = append(message.ImageURLs, a.EmbedItem.PlusPhoto.URL)
				if a.EmbedItem.PlusPhoto.URL == "https://lh3.googleusercontent.com/-kPexdUdQSuA/XQWkfNDso0I/AAAAAAAArYw/wmda3tF6vAIGfTWE9yJCzEi7KCsxyA2zgCK8BGAs/s0/2019-06-15.jpg" {
					fmt.Println("YEAH WE FOUND IT!!!")
				}
			}
			convo.Messages = append(convo.Messages, message)
		}
		if len(convo.Messages) < 1 {
			continue
		}
		chats = append(chats, convo)
		chatTitles = append(chatTitles, convo.Title)
		chatMap[convo.Title] = convo
	}

	prompt := &survey.Select{
		Message: "Choose a chat:",
		Options: chatTitles,
	}

	for {
		selectedChatTitle := ""
		err := survey.AskOne(prompt, &selectedChatTitle)
		if err != nil {
			break
		}

		selectedChat := chatMap[selectedChatTitle]

		filename := sanitize.Name(selectedChatTitle)

		c, err := chat.NewHTMLWriter("template.html", filename)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		err = c.WriteChat(selectedChat)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Printf("\r\n%v saved to %v\r\n\r\n", selectedChatTitle, filename)
	}

}
