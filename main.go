package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hussar_company/vstazherstve_back/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error

func CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var user models.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("abobus").Collection("dudes")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}

func main() {
	fmt.Println("Starting application")
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Hussar:Hussar1@hussarcluster.cokdm.mongodb.net/test"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(models.Herro())

	fmt.Println(databases)
	router := mux.NewRouter()

	router.HandleFunc("/user", CreateUserEndpoint).Methods("POST")

	http.ListenAndServe(":8000", router)
}
