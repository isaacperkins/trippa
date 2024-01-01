package controllers

import (
	users "cmd/api/internal/domain/users/model"
	"cmd/api/internal/domain/users/store"
	"encoding/json"
	"log"
	"net/http"
)

func UsersAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user users.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(store.AddUser(&user))
}
