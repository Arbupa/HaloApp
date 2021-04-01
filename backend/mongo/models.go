package mongo

import (
	"github.com/Arbupa/HaloApp/structs"
	"go.mongodb.org/mongo-driver/bson"
)

type MapsModel struct {
	ID                 bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name               string        `json:"name"`
	Description        string        `json:"description"`
	ImageUrl           string        `json:"imageUrl"`
	SupportedGameModes []string      `json:"supportedGameModes"`
}

type SkullsModel struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	IDApi       string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	ImageUrl    string        `json:"imageUrl"`
	// buscar esta madre
	MissionID string `json:"missionId"`
}

type MissionsModel struct {
	ID            bson.ObjectID `json:"id" bson:"_id,omitempty"`
	MissionNumber int           `json:"missionNumber"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	ImageUrl      string        `json:"imageUrl"`
	Type          string        `json:"type"`
}

type EnemiesModel struct {
	ID       bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Faction  string        `json:"faction"`
	Name     string        `json:"name"`
	ImageUrl string        `json:"largeIconImageUrl"`
}

type WeaponsModel struct {
	ID               bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	ImageUrl         string        `json:"largeIconImageUrl"`
	IsUsableByPlayer bool          `json:"isUsableByPlayer"`
	Type             string        `json:"type"`
}

type VehiclesModel struct {
	ID               bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Name             string        `json:"name"`
	Description      string        `json:"description"`
	ImageUrl         string        `json:"largeIconImageUrl"`
	IsUsableByPlayer bool          `json:"isUsableByPlayer"`
}

func MapToModel(mapp structs.Maps) MapsModel {
	var id bson.ObjectId
	if mapp.ID == "" {
		id = bson.NewObjectId()
	} else {
		id = bson.ObjectIdHex(mapp.ID)
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
