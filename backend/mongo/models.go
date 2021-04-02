package mongo

import (
	"github.com/Arbupa/HaloApp/structs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MapsModel struct {
	ID                 primitive.ObjectID `json:"_id" bson:"_id"`
	Name               string             `json:"name" bson:"name"`
	Description        string             `json:"description" bson:"description"`
	ImageUrl           string             `json:"image_url" bson:"image_url"`
	SupportedGameModes []string           `json:"supported_game_modes" bson:"supported_game_modes"`
}

type SkullsModel struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	IDApi       string             `json:"id" bson:"id_api"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	ImageUrl    string             `json:"image_url" bson:"image_url"`
	// buscar esta madre
	MissionID string `json:"mission_id" bson:"mission_id"`
}

type MissionsModel struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	MissionNumber int                `json:"mission_number" bson:"mission_number"`
	Name          string             `json:"name" bson:"name"`
	Description   string             `json:"description" bson:"description"`
	ImageUrl      string             `json:"image_url" bson:"image_url"`
	Type          string             `json:"type" bson:"type"`
}

type EnemiesModel struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Faction  string             `json:"faction" bson:"faction"`
	Name     string             `json:"name" bson:"name"`
	ImageUrl string             `json:"image_url" bson:"image_url"`
}

type WeaponsModel struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name             string             `json:"name" bson:"name"`
	Description      string             `json:"description" bson:"description"`
	ImageUrl         string             `json:"image_url" bson:"image_url"`
	IsUsableByPlayer bool               `json:"is_usable_by_player" bson:"is_usable_by_player"`
	Type             string             `json:"type" bson:"type"`
}

type VehiclesModel struct {
	ID               primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name             string             `json:"name" bson:"name"`
	Description      string             `json:"description" bson:"description"`
	ImageUrl         string             `json:"image_url" bson:"image_url"`
	IsUsableByPlayer bool               `json:"is_usable_by_player" bson:"is_usable_by_player"`
}

func MapsToModel(mapp []structs.Maps) []MapsModel {
	var MapsModels []MapsModel

	for i, _ := range mapp {

		var id primitive.ObjectID
		if mapp[i].ID == "" {
			id = primitive.NewObjectID()
		} else {
			id, _ = primitive.ObjectIDFromHex(mapp[i].ID)
		}

		mapModel := MapsModel{
			ID:                 id,
			Name:               mapp[i].Name,
			Description:        mapp[i].Description,
			ImageUrl:           mapp[i].ImageUrl,
			SupportedGameModes: mapp[i].SupportedGameModes,
		}
		MapsModels = append(MapsModels, mapModel)
	}

	return MapsModels
}

// en esta funcion es donde voy a buscar la mission por id
func SkullsToModel(skull []structs.Skulls) []SkullsModel {
	var models []SkullsModel

	for i, _ := range skull {

		var id primitive.ObjectID
		if skull[i].ID == "" {
			id = primitive.NewObjectID()
		} else {
			id, _ = primitive.ObjectIDFromHex(skull[i].ID)
		}

		skullModel := SkullsModel{
			ID:          id,
			IDApi:       skull[i].IDApi,
			Name:        skull[i].Name,
			Description: skull[i].Description,
			ImageUrl:    skull[i].ImageUrl,
			MissionID:   skull[i].MissionID,
			// Ojo acá
		}
		models = append(models, skullModel)
	}

	return models
}

func MissionsToModel(mission []structs.Missions) []MissionsModel {
	var models []MissionsModel

	for i, _ := range mission {

		var id primitive.ObjectID
		if mission[i].ID == "" {
			id = primitive.NewObjectID()
		} else {
			id, _ = primitive.ObjectIDFromHex(mission[i].ID)
		}
		missionModel := MissionsModel{
			ID:            id,
			MissionNumber: mission[i].MissionNumber,
			Name:          mission[i].Name,
			Description:   mission[i].Description,
			ImageUrl:      mission[i].ImageUrl,
			Type:          mission[i].Type,
		}
		models = append(models, missionModel)
	}

	return models
}

func EnemiesToModel(enemy []structs.Enemies) []EnemiesModel {
	var models []EnemiesModel

	for i, _ := range enemy {

		var id primitive.ObjectID
		if enemy[i].ID == "" {
			id = primitive.NewObjectID()
		} else {
			id, _ = primitive.ObjectIDFromHex(enemy[i].ID)
		}

		enemyModel := EnemiesModel{
			ID:       id,
			Faction:  enemy[i].Faction,
			Name:     enemy[i].Name,
			ImageUrl: enemy[i].ImageUrl,
		}
		models = append(models, enemyModel)
	}

	return models
}

func WeaponsToModel(weapon []structs.Weapons) []WeaponsModel {
	var models []WeaponsModel

	for i, _ := range weapon {

		var id primitive.ObjectID
		if weapon[i].ID == "" {
			id = primitive.NewObjectID()
		} else {
			id, _ = primitive.ObjectIDFromHex(weapon[i].ID)
		}

		weaponModel := WeaponsModel{
			ID:               id,
			Name:             weapon[i].Name,
			Description:      weapon[i].Description,
			ImageUrl:         weapon[i].ImageUrl,
			IsUsableByPlayer: weapon[i].IsUsableByPlayer,
			Type:             weapon[i].Type,
		}
		models = append(models, weaponModel)
	}

	return models
}

func VehiclesToModel(vehicle []structs.Vehicles) []VehiclesModel {
	var models []VehiclesModel

	for i, _ := range vehicle {

		var id primitive.ObjectID
		if vehicle[i].ID == "" {
			id = primitive.NewObjectID()
		} else {
			id, _ = primitive.ObjectIDFromHex(vehicle[i].ID)
		}

		vehicleModel := VehiclesModel{
			ID:               id,
			Name:             vehicle[i].Name,
			Description:      vehicle[i].Description,
			ImageUrl:         vehicle[i].ImageUrl,
			IsUsableByPlayer: vehicle[i].IsUsableByPlayer,
		}
		models = append(models, vehicleModel)
	}

	return models
}
