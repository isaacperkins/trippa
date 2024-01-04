package controllers

import (
	"net/http"
)

type Home struct {
	Title string `json:"title"`
	Copy  string `json:"copy"`
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	h := Home{
		Title: "velkum",
		Copy:  "asdf",
	}
	Result(w, h)
}
