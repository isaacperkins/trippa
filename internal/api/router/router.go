package router

import (
	"cmd/api/internal/api/controllers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(d *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.GetHome()).Methods("GET")
	r.HandleFunc("/users", controllers.UsersGet(d)).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UserGet(d)).Methods("GET")
	r.HandleFunc("/users", controllers.UsersAdd).Methods("PUT")
	r.HandleFunc("/users/{id}", controllers.UsersEdit).Methods("POST")
	r.HandleFunc("/users/{id}", controllers.UsersDelete).Methods("DELETE")

	return r
}
