package server

import (
	"time"
)

type Message struct {
	ID        int64
	Chat      *Chat
	Author    *User
	Text      string
	CreatedAt time.Time
}

type MessageService interface {
	Add(chat int64, author int64, text string) (int64, error)
}
