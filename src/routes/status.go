package routes

import (
	"net/http"
	"tool-checker/structures"
	"tool-checker/utils"
)

func Status(w http.ResponseWriter, r *http.Request) {

	response := structures.OK{"success", "it works!"}

	err := utils.Respond(w, response)
	if err != nil {
		panic(err)
	}

	return
}
