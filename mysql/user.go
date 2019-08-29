package mysql

import (
	"database/sql"

	server "github.com/dbond762/chat-server"
	_ "github.com/go-sql-driver/mysql"
)

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Add(username string) (int64, error) {
	result, err := us.DB.Exec(
		"INSERT INTO user (`username`) VALUES (?)",
		username,
	)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

func (us *UserService) Send(chat int64, author int64, text string) (int64, error) {
	return 0, nil
}

func (us *UserService) Chats(id int64) ([]server.Chat, error) {
	return []server.Chat{}, nil
}
