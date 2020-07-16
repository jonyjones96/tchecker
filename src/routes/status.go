package routes

import (
	"encoding/json"
	"net/http"
	"tool-checker/structures"
)

func Status(w http.ResponseWriter, r *http.Request) {

	response := structures.OK{"success", "it works!"}

	//err := json.NewEncoder(w).Encode("{'Status':'OK'}")
	//if err != nil {
	//panic(err)
	//}

	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)

	return
}
