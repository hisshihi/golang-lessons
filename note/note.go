package note

import (
	"errors"
	"time"
)

type Note struct {
	title     string
	content   string
	cteatedAt time.Time
}

func (note Note) Display() {
	println("Title:", note.title)
	println("Content:", note.content)
	println("Created At:", note.cteatedAt.Format(time.RFC1123))
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("input cannot be empty")
	}

	return Note{
		title:     title,
		content:   content,
		cteatedAt: time.Now(),
	}, nil
}
