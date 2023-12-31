package server

import (
	"cmd/api/internal/api/router"
	"cmd/api/internal/conf"
	db "cmd/api/internal/store/database"
	"fmt"
	"log"
	"net/http"
)

func Init() {

	c := conf.New()
	d := db.New(c)
	r := router.New(d)

	fmt.Println("Starting rest service: " + c.Http.Host + ":" + c.Http.Port)

	log.Fatal(http.ListenAndServe(c.Http.Host+":8080", r))

}
