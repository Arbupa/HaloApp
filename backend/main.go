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
	// Fill the database
	mongo.FillDatabase()
	// Create the router with his routes and methods
	r := mux.NewRouter()
	fmt.Println("Se ha inicializado el server")
	r.HandleFunc("/maps", http2.GetCampaignMissions).Methods("GET")

	//run the server
	log.Fatal(http.ListenAndServe(":5000", r))
}
