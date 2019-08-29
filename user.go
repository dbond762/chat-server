package server

import (
	"time"
)

type User struct {
	ID        int64
	Username  string
	CreatedAt time.Time
}

type UserService interface {
	Add(username string) (int64, error)
	Send(chat int64, author int64, text string) (int64, error)
	Chats(id int64) ([]Chat, error)
}
