package utils

import (
	"encoding/json"
	"net/http"
)

// Using practices recomended here: https://florimond.dev/blog/articles/2018/08/restful-api-design-13-best-practices-to-make-your-users-happy/
// Errors are values blog post: https://blog.golang.org/errors-are-values
func Respond(w http.ResponseWriter, response interface{}) error {
	jsonContent, err := json.Marshal(response)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonContent)

	return nil

}
