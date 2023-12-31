package main

import "cmd/api/internal/api/server"

/*/ Get /
func getBase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("base")
}
*/

func main() {
	server.Init()
}