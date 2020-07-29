package main

import (
	"log"
	"net/http"
	"tool-checker/routes"
	"tool-checker/utils"

	"github.com/gorilla/mux"
)

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (err error) {

	utils.ConnectDB()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/status", routes.Status).Methods("GET")
	router.HandleFunc("/api", routes.InsertTool).Methods("POST")
	router.HandleFunc("/tools", routes.GetTool).Methods("GET")

	log.Printf("API is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

	return nil
}

// Log handler discovered here: https://stackoverflow.com/questions/38443889/golang-logging-http-responses-in-addition-to-requests
// func logHandler(fn http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		x, err := httputil.DumpRequest(r, true)
// 		if err != nil {
// 			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
// 			return
// 		}
// 		log.Println(fmt.Sprintf("%q", x))
// 		rec := httptest.NewRecorder()
// 		fn(rec, r)
// 		log.Println(fmt.Sprintf("%q", rec.Body))
// 	}
// }
