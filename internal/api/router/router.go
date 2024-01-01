package router

import (
	"cmd/api/internal/domain/controllers"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	r.HandleFunc("/users", controllers.UsersGet).Methods("GET")
	r.HandleFunc("/users", controllers.UsersAdd).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.UsersEdit).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.UsersDelete).Methods("DELETE")
	return r
}
