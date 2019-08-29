package main

import (
	"database/sql"
	"log"

	"github.com/dbond762/chat-server/http"
	"github.com/dbond762/chat-server/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/chat_server?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	us := &mysql.UserService{DB: db}
	cs := &mysql.ChatService{DB: db}
	ms := &mysql.MessageService{DB: db}

	uh := &http.UserHandler{UserService: us}
	ch := &http.ChatHandler{ChatService: cs}
	mh := &http.MessageHandler{MessageService: ms}

	const port = 9000
	http.Setup(uh, ch, mh, port)
}
