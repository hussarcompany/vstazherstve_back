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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var user models.User
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("abobus").Collection("dudes")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}

func GetUsersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var users []models.User
	collection := client.Database("abobus").Collection("dudes")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message:"` + err.Error() + `"}`))
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message:"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(users)
}

func FindUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user models.User
	collection := client.Database("abobus").Collection("dudes")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, models.User{ID: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message:"` + err.Error() + `"}`))
	}
	json.NewEncoder(response).Encode(user)
}

func main() {
	fmt.Println("Starting application")
	var err error
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

	router.HandleFunc("/makeuser", CreateUserEndpoint).Methods("POST")
	router.HandleFunc("/users", GetUsersEndpoint).Methods("GET")
	router.HandleFunc("/finduser/{id}", FindUserEndpoint).Methods("GET")

	httpError := http.ListenAndServe(":8000", router)
	if httpError != nil {
		log.Println("While serving HTTP error: ", httpError)
	}
}
