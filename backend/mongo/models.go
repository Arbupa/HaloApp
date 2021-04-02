package mongo

import (
	"github.com/Arbupa/HaloApp/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type MapsModel struct {
	ID                 primitive.ObjectID `json:"_id" bson:"_id"`
	Name               string             `json:"name" bson:"name"`
	Description        string             `json:"description" bson:"description"`
	ImageUrl           string             `json:"image_url" bson:"image_url"`
	SupportedGameModes []string           `json:"supported_game_modes" bson:"supported_game_modes"`
}

type SkullsModel struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	IDApi       string        `json:"id" bson:"id_api"`
	Name        string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	ImageUrl    string        `json:"image_url" bson:"image_url"`
	// buscar esta madre
	MissionID string `json:"mission_id" bson:"mission_id"`
}

type MissionsModel struct {
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	MissionNumber int           `json:"mission_number" bson:"mission_number"`
	Name          string        `json:"name" bson:"name"`
	Description   string        `json:"description" bson:"description"`
	ImageUrl      string        `json:"image_url" bson:"image_url"`
	Type          string        `json:"type" bson:"type"`
}

type EnemiesModel struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Faction  string        `json:"faction" bson:"faction"`
	Name     string        `json:"name" bson:"name"`
	ImageUrl string        `json:"image_url" bson:"image_url"`
}

type WeaponsModel struct {
	ID               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name             string        `json:"name" bson:"name"`
	Description      string        `json:"description" bson:"description"`
	ImageUrl         string        `json:"image_url" bson:"image_url"`
	IsUsableByPlayer bool          `json:"is_usable_by_player" bson:"is_usable_by_player"`
	Type             string        `json:"type" bson:"type"`
}

type VehiclesModel struct {
	ID               bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name             string        `json:"name" bson:"name"`
	Description      string        `json:"description" bson:"description"`
	ImageUrl         string        `json:"image_url" bson:"image_url"`
	IsUsableByPlayer bool          `json:"is_usable_by_player" bson:"is_usable_by_player"`
}

func MapToModel(mapp structs.Maps) MapsModel {
	// for i, v := range mapp{

	// }
	var id primitive.ObjectID
	if mapp.ID == "" {
		id = primitive.NewObjectID()
	} else {
		//id = primitive.ObjectIDHex(mapp.ID)
		//id := id.(primitive.ObjectID).Hex()
	}

	mapModel := MapsModel{
		ID:                 id,
		Name:               mapp.Name,
		Description:        mapp.Description,
		ImageUrl:           mapp.ImageUrl,
		SupportedGameModes: mapp.SupportedGameModes,
	}

	return mapModel
}

// en esta funcion es donde voy a buscar la mission por id
func SkullToModel(skull structs.Skulls) SkullsModel {
	var id bson.ObjectId
	if skull.ID == "" {
		id = bson.NewObjectId()
	} else {
		id = bson.ObjectIdHex(skull.ID)
	}

	skullModel := SkullsModel{
		ID:          id,
		IDApi:       skull.IDApi,
		Name:        skull.Name,
		Description: skull.Description,
		ImageUrl:    skull.ImageUrl,
		// Ojo acá
		MissionID: skull.MissionID,
	}

	return skullModel
}

func MissionToModel(mission structs.Missions) MissionsModel {
	var id bson.ObjectId
	if mission.ID == "" {
		id = bson.NewObjectId()
	} else {
		id = bson.ObjectIdHex(mission.ID)
	}

	missionModel := MissionsModel{
		ID:            id,
		MissionNumber: mission.MissionNumber,
		Name:          mission.Name,
		Description:   mission.Description,
		ImageUrl:      mission.ImageUrl,
		Type:          mission.Type,
	}

	return missionModel
}

func EnemiesToModel(enemy structs.Enemies) EnemiesModel {
	var id bson.ObjectId
	if enemy.ID == "" {
		id = bson.NewObjectId()
	} else {
		id = bson.ObjectIdHex(enemy.ID)
	}

	enemyModel := EnemiesModel{
		ID:       id,
		Faction:  enemy.Faction,
		Name:     enemy.Name,
		ImageUrl: enemy.ImageUrl,
	}

	return enemyModel
}

func WeaponsToModel(weapon structs.Weapons) WeaponsModel {
	var id bson.ObjectId
	if weapon.ID == "" {
		id = bson.NewObjectId()
	} else {
		id = bson.ObjectIdHex(weapon.ID)
	}

	weaponModel := WeaponsModel{
		ID:               id,
		Name:             weapon.Name,
		Description:      weapon.Description,
		ImageUrl:         weapon.ImageUrl,
		IsUsableByPlayer: weapon.IsUsableByPlayer,
		Type:             weapon.Type,
	}

	return weaponModel
}

func VehiclesToModel(vehicle structs.Vehicles) VehiclesModel {
	var id bson.ObjectId
	if vehicle.ID == "" {
		id = bson.NewObjectId()
	} else {
		id = bson.ObjectIdHex(vehicle.ID)
	}

	vehicleModel := VehiclesModel{
		ID:               id,
		Name:             vehicle.Name,
		Description:      vehicle.Description,
		ImageUrl:         vehicle.ImageUrl,
		IsUsableByPlayer: vehicle.IsUsableByPlayer,
	}

	return vehicleModel
}
