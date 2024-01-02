package controllers

import (
	users "cmd/api/internal/domain/users/model"
	"cmd/api/internal/domain/users/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func UsersGet(d *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Result(w, store.UsersGet(d))
	}
}

func UserGet(d *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		Result(w, store.UserGet(d, mux.Vars(r)["id"]))
	}
}

func UsersAdd(w http.ResponseWriter, r *http.Request) {
	var user users.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	Result(w, store.UsersAdd(&user))
}

func UsersEdit(w http.ResponseWriter, r *http.Request) {
	var user users.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	Result(w, store.UsersEdit(&user))
}

func UsersDelete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	Result(w, store.UsersDelete(id))
}
