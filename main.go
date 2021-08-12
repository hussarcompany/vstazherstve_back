package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func main() {
	fmt.Println("Starting application")

	var err error
	Client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Hussar:Hussar1@hussarcluster.cokdm.mongodb.net/test"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = Client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer Client.Disconnect(ctx)
	databases, err := Client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)

	router := mux.NewRouter()

	router.HandleFunc("/makeuser", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/users", GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/finduser/{id}", FindUserEndpoint).Methods("GET")

	httpError := http.ListenAndServe(":8000", router)
	if httpError != nil {
		log.Println("While serving HTTP error: ", httpError)
	}
}
