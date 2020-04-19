package main

import (
  "fmt"
  "log"
	"net/http"
  "context"
  "time"

  "github.com/gorilla/mux"
  "tool-checker/routes"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/readpref"
  "go.mongodb.org/mongo-driver/mongo/options"

)

type Person struct {
    Name string
    Age  int
    City string
}


func main() {

  fmt.Println("STarting the application")
  err := run()
  if err != nil {
    log.Fatal(err)
  }

}

func run() (err error) {

  connectMongo()  

	router := mux.NewRouter().StrictSlash(true)
  //	router.HandleFunc("/", homeLink)
	router.HandleFunc("/api", routes.InsertTool).Methods("POST")

  fmt.Printf("API is running on port 8080...")
  log.Fatal(http.ListenAndServe(":8080",router))

  return nil
}

func connectMongo() {

   var (
       client     *mongo.Client
       mongoURL = "mongodb://mongo:27017"
   )

   // Initialize a new mongo client with options
   client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL).SetAuth(options.Credential{Username: "user", Password: "example",}))

   // Connect the mongo client to the MongoDB server
   ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
   err = client.Connect(ctx)

   // Ping MongoDB
   ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
   if err = client.Ping(ctx, readpref.Primary()); err != nil {
       fmt.Println("could not ping to mongo db service: %v\n", err)
       return
   }

   fmt.Println("connected to nosql database:", mongoURL)

  // Insert test
  collection := client.Database("mydb").Collection("persons")

    ruan := Person{"Ruan", 34, "Cape Town"}

    insertResult, err := collection.InsertOne(context.TODO(), ruan)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)

}


