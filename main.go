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
	"github.com/vivanshah/hangout-json-parser/chatWriter"
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

	json.Unmarshal(byteValue, &hangouts)

	fmt.Println("Loaded ", len(hangouts.Conversations), " conversations")
	chatMap := map[string]models.Chat{}
	chatTitles := []string{}
	chats := []models.Chat{}
	for _, c := range hangouts.Conversations {
		chat := models.Chat{
			ParticipantIDs:   map[string]int{},
			ParticipantNames: map[string]string{},
			Messages:         []models.Message{},
		}

		participants := c.Conversation.Conversation.ParticipantData
		for n, p := range participants {
			chat.ParticipantIDs[p.ID.ChatID] = n
			chat.ParticipantNames[p.ID.ChatID] = p.FallbackName
			chat.Title = chat.Title + p.FallbackName + ", "
		}
		chat.Title = strings.TrimRight(chat.Title, ", ")
		for _, e := range c.Events {
			if e.EventType != "REGULAR_CHAT_MESSAGE" {
				continue
			}
			message := models.Message{}
			t, _ := strconv.ParseInt(e.Timestamp, 10, 64)
			t = t * 1000
			message.Timestamp = time.Unix(0, t)
			message.SenderID = chat.ParticipantIDs[e.SenderID.ChatID]
			message.Sender = chat.ParticipantNames[e.SenderID.ChatID]
			for _, s := range e.ChatMessage.MessageContent.Segment {
				message.Text = message.Text + " " + s.Text
				if s.Type == "LINK" {
					message.Text = message.Text + "[" + s.LinkData.LinkTarget + "]"
				}
			}
			chat.Messages = append(chat.Messages, message)
		}
		chats = append(chats, chat)
		chatTitles = append(chatTitles, chat.Title)
		chatMap[chat.Title] = chat
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

		filename := sanitize.Name(selectedChatTitle + ".txt")

		c, err := chatWriter.NewTxtWriter(filename)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		err = c.WriteChat(selectedChat)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(selectedChatTitle, " saved to ", filename)
	}

}
