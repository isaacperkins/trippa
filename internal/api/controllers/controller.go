package controllers

import (
	"encoding/json"
	"net/http"
)

func Result(w http.ResponseWriter, o interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(o)
}
