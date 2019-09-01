package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/dbond762/chat-server/http"
	"github.com/dbond762/chat-server/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		port       = flag.Int("port", 9000, "")
		dbUser     = flag.String("db-user", "root", "")
		dbPassword = flag.String("db-password", "", "")
		dbHost     = flag.String("db-host", "localhost", "")
		dbPort     = flag.Int("db-port", 3306, "")
		dbName     = flag.String("db-name", "chat_server", "")
	)

	flag.Parse()

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", *dbUser, *dbPassword, *dbHost, *dbPort, *dbName)
	db, err := sql.Open("mysql", dataSource)
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

	http.Setup(uh, ch, mh, *port)
}
