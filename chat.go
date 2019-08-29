package server

import (
	"time"
)

type Chat struct {
	ID        int64
	Name      string
	Users     []User
	CreatedAt time.Time
}

type ChatService interface {
	Add(name string, users []int64) (int64, error)
	Messages(id int64) ([]Message, error)
}
