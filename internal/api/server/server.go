package server

import (
	"cmd/api/internal/api/conf"
	"cmd/api/internal/api/db"
	"cmd/api/internal/api/router"
	"fmt"
	"log"
	"net/http"
)

func Init() {
	c := conf.New("conf/config.yaml")
	fmt.Println("Port:" + c.Http.Port)

	db.Init()
	db.AddUser("test")

	r := router.New()

	log.Fatal(http.ListenAndServe(":8080", r))
}
