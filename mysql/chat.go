package mysql

import (
	"database/sql"

	server "github.com/dbond762/chat-server"
)

type ChatService struct {
	DB *sql.DB
}

func (cs *ChatService) Add(name string, users []int64) (int64, error) {
	tx, err := cs.DB.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	result, err := tx.Exec(
		"INSERT INTO chat (`name`) VALUES (?)",
		name,
	)
	if err != nil {
		return 0, err
	}

	chatID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	stmt, err := tx.Prepare("INSERT INTO user_chat (`user_id`, `chat_id`) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	for _, userID := range users {
		_, err := stmt.Exec(userID, chatID)
		if err != nil {
			return 0, err
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return chatID, nil
}

func (ch *ChatService) Messages(id int64) ([]server.Message, error) {
	return []server.Message{}, nil
}
