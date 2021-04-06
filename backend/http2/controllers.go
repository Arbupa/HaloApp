package http2

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

var Database = Db().Database("HaloApp")

func GetAllMaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []primitive.M

	missionCollection := Database.Collection("maps") //returns a *mongo.Cursor
	cur, err := missionCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}

func GetAllSkulls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []primitive.M

	missionCollection := Database.Collection("skulls") //returns a *mongo.Cursor
	cur, err := missionCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}

func GetAllMissions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []primitive.M

	missionCollection := Database.Collection("missions") //returns a *mongo.Cursor
	cur, err := missionCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}

func GetAllEnemies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []primitive.M

	missionCollection := Database.Collection("enemies") //returns a *mongo.Cursor
	cur, err := missionCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}

func GetAllWeapons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []primitive.M

	missionCollection := Database.Collection("weapons") //returns a *mongo.Cursor
	cur, err := missionCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}

func GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var results []primitive.M

	missionCollection := Database.Collection("vehicles") //returns a *mongo.Cursor
	cur, err := missionCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var elem primitive.M
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem) // appending document pointed by Next()
	}
	cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	json.NewEncoder(w).Encode(results)
}
