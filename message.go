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
