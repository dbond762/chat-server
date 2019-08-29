package http

import (
	"encoding/json"
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
