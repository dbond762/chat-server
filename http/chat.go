package http

import (
	"encoding/json"
	"log"
	"net/http"

	server "github.com/dbond762/chat-server"
)

type ChatHandler struct {
	ChatService server.ChatService
}

type ChatAddRequest struct {
	Name  string  `json:"name"`
	Users []int64 `json:"users"`
}

func (ch *ChatHandler) Add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := new(ChatAddRequest)
	if err := decoder.Decode(request); err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := ch.ChatService.Add(request.Name, request.Users)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	response := &Response{ID: id}
	if err := encoder.Encode(response); err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

type ChatMessagesRequset struct {
	Chat int64 `json:"chat"`
}

type MessageItem struct {
	ID        int64    `json:"id"`
	Chat      int64    `json:"chat"`
	Author    UserItem `json:"author"`
	Text      string   `json:"text"`
	CreatedAt string   `json:"created_at"`
}

type ChatMessagesResponse struct {
	Messages []MessageItem `json:"messages"`
}

func (ch *ChatHandler) Messages(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := new(ChatMessagesRequset)
	if err := decoder.Decode(request); err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	messages, err := ch.ChatService.Messages(request.Chat)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)

	response := new(ChatMessagesResponse)
	for _, message := range messages {
		messageItem := MessageItem{
			ID:   message.ID,
			Chat: message.Chat,
			Author: UserItem{
				ID:        message.Author.ID,
				Username:  message.Author.Username,
				CreatedAt: message.Author.CreatedAt,
			},
			Text:      message.Text,
			CreatedAt: message.CreatedAt,
		}
		response.Messages = append(response.Messages, messageItem)
	}

	if err := encoder.Encode(response); err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
