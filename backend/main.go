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
	// fix this thing to not always fill Database
	// data := mongo.DataExists()
	// fmt.Println(data)
	// if data == false {
	// 	// Fill the database
	// 	mongo.FillDatabase()
	// }
	mongo.FillDatabase()
	// Create the router with his routes and methods
	r := mux.NewRouter()
	fmt.Println("Se ha inicializado el server")

	// Routes
	r.HandleFunc("/maps", http2.GetAllMaps).Methods("GET")
	r.HandleFunc("/missions", http2.GetAllMissions).Methods("GET")
	r.HandleFunc("/skulls", http2.GetAllSkulls).Methods("GET")
	r.HandleFunc("/enemies", http2.GetAllEnemies).Methods("GET")
	r.HandleFunc("/weapons", http2.GetAllWeapons).Methods("GET")
	r.HandleFunc("/vehicles", http2.GetAllVehicles).Methods("GET")

	//run the server
	log.Fatal(http.ListenAndServe(":5000", r))
}
