package server

type Message struct {
	ID        int64
	Chat      int64 //*Chat
	Author    int64 //*User
	Text      string
	CreatedAt string
}

type MessageService interface {
	Add(chat int64, author int64, text string) (int64, error)
}
