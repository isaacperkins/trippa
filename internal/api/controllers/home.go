package controllers

import (
	"net/http"
)

type home struct {
	title string
	copy  string
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	h := home{title: "velkum", copy: "asdf"}
	Result(w, h)
}
