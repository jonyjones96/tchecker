package routes

import (
	"encoding/json"
	"net/http"
)

func Status(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode("{'Status':'OK'}")

	return
}