package mongo

import (
	"context"
	"fmt"
	"time"
)

var Database = Db().Database("HaloApp") // get collection "HaloApp" from Db() which returns *mongo.Client

func InsertMaps(arr MapsModel) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	//defer client.Disconnect(ctx)
	mapsCollection := Database.Collection("maps")

	//for _, v := range arr {
	insertMap, err := mapsCollection.InsertOne(ctx, arr)
	if err != nil {
		fmt.Println(err)
	}
	if insertMap != nil {
		fmt.Println(insertMap.InsertedID)
	}
	//}
}
