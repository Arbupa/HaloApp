package mongo

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertMaps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client, err := mongo.NewClient(options.Client())
	if err != nil {
		log.Fatal(err)
	}
	database := client.Database("Halo")
	defer client.Disconnect(context.TODO())
	mapsCollection := database.Collection("maps")

}
