package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tool-checker/structures"
	"tool-checker/utils"
)

func InsertTool(w http.ResponseWriter, r *http.Request) {

	// Insert test
	collection := utils.Client.Database("mydb").Collection("tools")

	newTool := structures.Tool{}

	err := json.NewDecoder(r.Body).Decode(&newTool)
	if err != nil {
		log.Fatal(err)
	}

	insertResult, err := collection.InsertOne(context.TODO(), newTool)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

	return
}
