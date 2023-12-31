package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("base")
}
