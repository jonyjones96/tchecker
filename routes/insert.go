package routes

import (
  "fmt"
  "net/http"
)

func InsertTool(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome Home!")
}
