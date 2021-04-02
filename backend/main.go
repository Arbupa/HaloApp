package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Arbupa/HaloApp/http2"
	"github.com/Arbupa/HaloApp/mongo"
	"github.com/gorilla/mux"
)

func main() {

	// Get all the data from Halo Api
	mapsArrays := http2.GetMaps()
	fmt.Println(mapsArrays[0])
	// convert the original struct to mongo model
	mapModel := mongo.MapToModel(mapsArrays[0])
	fmt.Println("EL MODELO DE MONGO ES: ")
	fmt.Println(mapModel)
	// Create the connection with mongo
	DbConnection := mongo.Db()
	fmt.Println(DbConnection)
	// Insert the previous data inside the database
	mongo.InsertMaps(mapModel)
	// Create the router with his routes and methods
	r := mux.NewRouter()
	fmt.Println("Se ha inicializado el server")
	r.HandleFunc("/maps", http2.GetCampaignMissions).Methods("GET")

	//run the server
	log.Fatal(http.ListenAndServe(":5000", r))
}
