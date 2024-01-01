package server

import (
	"cmd/api/internal/api/router"
	"cmd/api/internal/conf"
	"cmd/api/internal/store/pgsql"
	"fmt"
	"log"
	"net/http"
)

func Init() {
	c := conf.New()
	fmt.Println("Port:" + c.Http.Port)

	db.Init()

	r := router.New()

	log.Fatal(http.ListenAndServe(c.Http.Host+":8080", r))
}
