package controllers

import (
	users "cmd/api/internal/domain/users/model"
	"cmd/api/internal/domain/users/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func UsersGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(store.UsersGet())
}

func UsersAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user users.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(store.UsersAdd(&user))
}

func UsersEdit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user users.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}
	json.NewEncoder(w).Encode(store.UsersEdit(&user))
}

func UsersDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	json.NewEncoder(w).Encode(store.UsersDelete(id))
}
