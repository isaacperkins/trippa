package router

import (
	"cmd/api/internal/api/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(d *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	r.HandleFunc("/users", controllers.UsersGet(d)).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UserGet(d)).Methods("GET")
	r.HandleFunc("/users", controllers.UsersAdd(d)).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.UsersEdit(d)).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.UsersDelete(d)).Methods("DELETE")

	return r
}
