package controllers

import (
	"net/http"
)

type home struct {
	title string
	copy  string
}

func GetHome() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//store.Migrate(d)
		h := home{title: "velkum", copy: "asdf"}
		Result(w, h)
	}
}
