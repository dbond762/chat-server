package http

import (
	"encoding/json"
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
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := ch.ChatService.Add(request.Name, request.Users)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)
	response := &Response{ID: id}
	if err := encoder.Encode(response); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
