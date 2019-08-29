package http

import (
	"encoding/json"
	"net/http"

	server "github.com/dbond762/chat-server"
)

type MessageHandler struct {
	MessageService server.MessageService
}

type MessageAddRequest struct {
	Chat   int64  `json:"chat"`
	Author int64  `json:"author"`
	Text   string `json:"text"`
}

func (mh *MessageHandler) Add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := new(MessageAddRequest)
	if err := decoder.Decode(request); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := mh.MessageService.Add(request.Chat, request.Author, request.Text)
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
