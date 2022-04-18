package service

import (
	"context"
	"crud/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"crud/models"
)

var collection *mongo.Collection = db.ConnectDb()

//Creating a new product according to the client request.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)

	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result)
}

//Getting all the products
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var users []models.User
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

//Updating the user according to the client request.
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	id := r.URL.Query().Get("id")
	fmt.Println("id =>", id)

	ids, _ := primitive.ObjectIDFromHex(id)

	var user models.User

	filter := bson.M{"_id": ids}

	_ = json.NewDecoder(r.Body).Decode(&user)

	update := bson.D{
		{"$set", bson.D{
			{"name", user.Name},
			{"gender", user.Gender},
			{"age", user.Age},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	user.ID = ids

	json.NewEncoder(w).Encode(user)

}


//Deleting the user according to the client request.
func DeleteUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	fmt.Println("id =>", id)

	ids, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err, w)
	}
	fmt.Println(ids)

	filter := bson.M{"_id": ids}

	

	var users models.User
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		users = user
	}


	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err, w)
	}

	fmt.Println(deleteResult)

	w.WriteHeader(http.StatusAccepted)

	json.NewEncoder(w).Encode(users)
}