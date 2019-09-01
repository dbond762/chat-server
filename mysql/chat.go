package mysql

import (
	"database/sql"

	server "github.com/dbond762/chat-server"
	_ "github.com/go-sql-driver/mysql"
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

func (cs *ChatService) Messages(id int64) ([]server.Message, error) {
	rows, err := cs.DB.Query(
		"SELECT `message`.`id` AS `id`, `message`.`chat_id`, `user`.`id` AS `author_id`, `user`.`username` AS `author_username`, `user`.`created_at` AS `author_created_at`, `message`.`text` AS `text`, `message`.`created_at` AS `created_at`"+
			"FROM `message` LEFT JOIN `user` ON `message`.`author_id` = `user`.`id`"+
			"WHERE `message`.`chat_id` = ? "+
			"ORDER BY `created_at` DESC",
		id,
	)
	if err != nil {
		return []server.Message{}, err
	}

	messages := make([]server.Message, 0)

	for rows.Next() {
		var (
			m server.Message
			a server.User
		)

		if err := rows.Scan(&m.ID, &m.Chat, &a.ID, &a.Username, &a.CreatedAt, &m.Text, &m.CreatedAt); err != nil {
			return []server.Message{}, err
		}

		m.Author = &a
		messages = append(messages, m)
	}
	rows.Close()

	return messages, nil
}
