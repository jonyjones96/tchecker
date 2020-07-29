package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tool-checker/structures"
	"tool-checker/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertTool adds a new tool to the database
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

	message := structures.OK{"Success", fmt.Sprintf("Inserted a Single Document: %s", insertResult.InsertedID)}

	utils.Respond(w, message)

	return
}

// GetTool ...
func GetTool(w http.ResponseWriter, r *http.Request) {

	// Information found here: https://www.mongodb.com/blog/post/mongodb-go-driver-tutorial

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(5)

	collection := utils.Client.Database("mydb").Collection("tools")

	var results []*structures.Tool

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem structures.Tool
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	// return the list to the user
	utils.Respond(w, results)
}
