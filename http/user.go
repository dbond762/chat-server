package http

import (
	"encoding/json"
	"log"
	"net/http"

	server "github.com/dbond762/chat-server"
)

type UserHandler struct {
	UserService server.UserService
}

type UserAddRequest struct {
	Username string `json:"username"`
}

func (uh *UserHandler) Add(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := new(UserAddRequest)
	if err := decoder.Decode(request); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := uh.UserService.Add(request.Username)
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

type UserChatsRequest struct {
	User int64 `json:"user"`
}

type ChatItem struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `jsog:"created_at"`
}

// todo add users
// todo add sort
type UserChatsResponse struct {
	Chats []ChatItem `json:"chats"`
}

func (uh *UserHandler) Chats(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := new(UserChatsRequest)
	if err := decoder.Decode(request); err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	chats, err := uh.UserService.Chats(request.User)
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(w)

	response := new(UserChatsResponse)
	for _, chat := range chats {
		chatItem := ChatItem{
			ID:        chat.ID,
			Name:      chat.Name,
			CreatedAt: chat.CreatedAt,
		}
		response.Chats = append(response.Chats, chatItem)
	}

	if err := encoder.Encode(response); err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
