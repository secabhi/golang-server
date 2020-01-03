package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Username string `string:"username"`
	Email    string `string:"email"`
	Password string `string:"password"`
}

var client *mongo.Client

func dbConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(
		// "<put mongo path>",
	))
}

func SignUpUser(w http.ResponseWriter, t *http.Request) {
	w.Header().Add("content-type", "application/json")
	var user User
	err := json.NewDecoder(t.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	collection := client.Database("development").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, _ := collection.InsertOne(ctx, user)
	// fmt.Fprintf(w, "Person: %+v", res)
	json.NewEncoder(w).Encode(res)
}

func LoginUser(w http.ResponseWriter, t *http.Request) {
	fmt.Println("LoginUser---")
}

func AddUser(w http.ResponseWriter, t *http.Request) {
	fmt.Println("AddUser---")
}

func RemoveUser(w http.ResponseWriter, t *http.Request) {
	fmt.Println("removeUser---")
}

func UpdateUser(w http.ResponseWriter, t *http.Request) {
	fmt.Println("updateUser---")
}

func AllUser(w http.ResponseWriter, t *http.Request) {
	w.Header().Add("content-type", "application/json")
	var users []User
	collection := client.Database("development").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var usr User
		cursor.Decode(&usr)
		users = append(users, usr)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	// fmt.Fprintf(w, "Person: %+v", res)
	json.NewEncoder(w).Encode(users)
}
