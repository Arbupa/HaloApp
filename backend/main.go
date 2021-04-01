package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/arbupa/HaloApp/http2"
	"gitlab.com/arbupa/HaloApp/mongo"
)

func main() {
	// Create the connection with mongo
	DbConnection := mongo.Db()
	fmt.Println(DbConnection)

	// Get all the data from Halo Api

	// Create the router with his routes and methods
	r := mux.NewRouter()
	r.HandleFunc("/", http2.GetCampaignMissions).Methods("GET")

	//run the server
	log.Fatal(http.ListenAndServe(":5000", r))
}
