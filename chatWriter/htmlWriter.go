package chat

import (
	"html/template"
	"os"

	"github.com/vivanshah/hangout-json-parser/models"
)

type HTMLWriter struct {
	Path string
	tmpl *template.Template
}

func NewHTMLWriter(templatePath string, path string) (Writer, error) {
	t := template.Must(template.ParseFiles(templatePath))
	h := HTMLWriter{
		Path: path + ".html",
		tmpl: t,
	}

	return &h, nil
}

func (t *HTMLWriter) WriteChat(chat models.Chat) error {

	f, err := os.Create(t.Path)
	defer f.Close()
	if err != nil {
		return err
	}

	err = t.tmpl.Execute(f, chat)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
