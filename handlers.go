package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/hussar_company/vstazherstve_back/models/client"
	"github.com/hussar_company/vstazherstve_back/models/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtKey = []byte("secret_key")

func Login(response http.ResponseWriter, request *http.Request) {
	var user client.User
	print(user)

}

func CreateUserEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var user dto.UserDTO
	_ = json.NewDecoder(request.Body).Decode(&user)
	collection := Client.Database("abobus").Collection("dudes")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, user)
	json.NewEncoder(response).Encode(result)
}

func GetUsersEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var users []dto.UserDTO
	collection := Client.Database("abobus").Collection("dudes")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message:"` + err.Error() + `"}`))
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user dto.UserDTO
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
	var user dto.UserDTO
	collection := Client.Database("abobus").Collection("dudes")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, dto.UserDTO{ID: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message:"` + err.Error() + `"}`))
	}
	json.NewEncoder(response).Encode(user)
}
