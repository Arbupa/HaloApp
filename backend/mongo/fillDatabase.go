package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/Arbupa/HaloApp/http2"
)

var Database = Db().Database("HaloApp") // get collection "HaloApp" from Db() which returns *mongo.Client

func InsertMaps(arr []MapsModel) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	for _, v := range arr {

		//defer client.Disconnect(ctx)
		mapsCollection := Database.Collection("maps")

		insertMap, err := mapsCollection.InsertOne(ctx, v)
		if err != nil {
			fmt.Println(err)
		}
		if insertMap != nil {
			fmt.Println(insertMap.InsertedID)
		}
	}
}

func InsertEnemies(arr []EnemiesModel) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	for _, v := range arr {

		//defer client.Disconnect(ctx)
		mapsCollection := Database.Collection("enemies")

		insertEnemy, err := mapsCollection.InsertOne(ctx, v)
		if err != nil {
			fmt.Println(err)
		}
		if insertEnemy != nil {
			fmt.Println(insertEnemy.InsertedID)
		}
	}
}

func InsertSkulls(arr []SkullsModel) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	for _, v := range arr {

		skullsCollection := Database.Collection("skulls")
		insertSkull, err := skullsCollection.InsertOne(ctx, v)
		if err != nil {
			fmt.Println(err)
		}
		if insertSkull != nil {
			fmt.Println(insertSkull.InsertedID)
		}
	}
}

func InsertMissions(arr []MissionsModel) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	for _, v := range arr {

		missionsCollection := Database.Collection("missions")
		insertMission, err := missionsCollection.InsertOne(ctx, v)
		if err != nil {
			fmt.Println(err)
		}
		if insertMission != nil {
			fmt.Println(insertMission.InsertedID)
		}
	}
}

func InsertWeapons(arr []WeaponsModel) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	for _, v := range arr {

		weaponsCollection := Database.Collection("weapons")
		insertWeapon, err := weaponsCollection.InsertOne(ctx, v)
		if err != nil {
			fmt.Println(err)
		}
		if insertWeapon != nil {
			fmt.Println(insertWeapon.InsertedID)
		}
	}
}

func InsertVehicles(arr []VehiclesModel) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	for _, v := range arr {

		vehiclesCollection := Database.Collection("vehicles")
		insertVehicle, err := vehiclesCollection.InsertOne(ctx, v)
		if err != nil {
			fmt.Println(err)
		}
		if insertVehicle != nil {
			fmt.Println(insertVehicle.InsertedID)
		}
	}
}

func FillDatabase() error {
	// maps operations
	mapsArrays := http2.GetMaps()
	mapModels := MapsToModel(mapsArrays)
	InsertMaps(mapModels)
	// ahora las de skulls
	skullsArrays := http2.GetSkulls()
	skullsModels := SkullsToModel(skullsArrays)
	InsertSkulls(skullsModels)
	// ahora Missiones
	missionsArrays := http2.GetMissions()
	missionsModels := MissionsToModel(missionsArrays)
	InsertMissions(missionsModels)
	// ahora Enemies
	enemiesArrays := http2.GetEnemies()
	enemiesModels := EnemiesToModel(enemiesArrays)
	InsertEnemies(enemiesModels)
	// ahora Weapons
	weaponsArrays := http2.GetWeapons()
	weaponsModels := WeaponsToModel(weaponsArrays)
	InsertWeapons(weaponsModels)
	// ahora Vehicles
	vehiclesArrays := http2.GetVehicles()
	vehiclesModels := VehiclesToModel(vehiclesArrays)
	InsertVehicles(vehiclesModels)
	return nil
}
