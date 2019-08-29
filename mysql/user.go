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

func (us *UserService) Chats(id int64) ([]server.Chat, error) {
	rows, err := us.DB.Query(
		"SELECT `chat`.`id`, `chat`.`name`, `chat`.`created_at` FROM `chat` LEFT JOIN `user_chat` ON `chat`.`id` = `user_chat`.`chat_id` WHERE `user_chat`.`user_id` = ?",
		id,
	)
	if err != nil {
		return []server.Chat{}, err
	}

	const defaultCapacity = 25
	chats := make([]server.Chat, 0, defaultCapacity)

	for rows.Next() {
		var c server.Chat
		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt); err != nil {
			return []server.Chat{}, err
		}
		chats = append(chats, c)
	}
	rows.Close()

	return chats, nil
}
