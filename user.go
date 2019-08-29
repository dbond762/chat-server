package server

type User struct {
	ID        int64
	Username  string
	CreatedAt string
}

type UserService interface {
	Add(username string) (int64, error)
	Chats(id int64) ([]Chat, error)
}
