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

func UsersAdd(d *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user users.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println(err)
		}

		Result(w, store.UsersAdd(d, &user))
	}
}

func UsersEdit(d *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user users.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Println(err)
		}

		Result(w, store.UsersEdit(d, &user))
	}
}

func UsersDelete(d *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		Result(w, store.UsersDelete(d, id))
	}
}
