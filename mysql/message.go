package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MessageService struct {
	DB *sql.DB
}

func (ms *MessageService) Add(chat int64, author int64, text string) (int64, error) {
	result, err := ms.DB.Exec(
		"INSERT INTO message (`chat_id`, `author_id`, `text`) VALUES (?, ?, ?)",
		chat,
		author,
		text,
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
