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
		"SELECT `t1`.`id` AS `chat_id`, `t1`.`name` AS `chat_name`, `t1`.`created_at` AS `chat_created_at`, `user`.`id` AS `user_id`, `user`.`username` AS `user_username`, `user`.`created_at` AS `user_created_at`"+
			"FROM"+
			"	("+
			"		SELECT `chat`.`id` AS `id`, `chat`.`name` AS `name`, `chat`.`created_at` AS `created_at`"+
			"		FROM `chat` LEFT JOIN `user_chat` ON `chat`.`id` = `user_chat`.`chat_id`"+
			"		WHERE `user_chat`.`user_id` = ?"+
			"	) AS `t1`"+
			"	LEFT JOIN `user_chat` ON `t1`.`id` = `user_chat`.`chat_id`"+
			"	LEFT JOIN `user` ON `user_chat`.`user_id` = `user`.`id`"+
			"ORDER BY `chat_created_at` DESC, `user_created_at` DESC",
		id,
	)
	if err != nil {
		return []server.Chat{}, err
	}

	chats := make([]server.Chat, 0)

	chatIdx := 0
	for rows.Next() {
		var c server.Chat
		var u server.User

		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt, &u.ID, &u.Username, &u.CreatedAt); err != nil {
			return []server.Chat{}, err
		}

		if chatIdx != 0 && c.ID == chats[chatIdx-1].ID {
			chats[chatIdx-1].Users = append(chats[chatIdx-1].Users, u)
		} else {
			c.Users = append(c.Users, u)
			chats = append(chats, c)
			chatIdx++
		}
	}
	rows.Close()

	return chats, nil
}
