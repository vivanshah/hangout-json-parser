package chat

import (
	"fmt"
	"os"
)

type txtWriter struct {
	Path string
}

func NewTxtWriter(filename string) (Writer, error) {
	t := txtWriter{
		Path: filename + ".txt",
	}

	return &t, nil
}

func (t *txtWriter) WriteChat(convo Conversation) error {

	f, err := os.Create(t.Path)
	defer f.Close()
	if err != nil {
		return err
	}

	var line string
	for _, m := range convo.Messages {
		line = m.Sender + " @ " + m.Timestamp.Format("3:04:05 PM Jan _2 2006") + ": " + m.Text + "\r"
		fmt.Fprintln(f, line)
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
