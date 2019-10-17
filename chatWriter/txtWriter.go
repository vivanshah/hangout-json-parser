package chatWriter

import (
	"os"
	"fmt"
	"github.com/vivanshah/hangout-json-parser/models"
)

type txtWriter struct {
	Path string
}

func NewTxtWriter(path string) (ChatWriter, error) {
	t := txtWriter {
		Path: path,
	}

	return &t, nil
}



func (t *txtWriter)  WriteChat(chat models.Chat) error {

	f, err := os.Create(t.Path)
	defer f.Close()
	if err != nil {
		return err
	}

	var line string
	for _, m := range chat.Messages {
		line = m.Sender + " @ " + m.Timestamp.Format("3:04:05 PM Jan _2 2006") + ": " + m.Text + "\r"
		fmt.Fprintln(f, line)
	}

	

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}