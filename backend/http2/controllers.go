package http2

//Agregar una interface para usar esas funciones/objetos que necesito y evitar el loop

import (
	"net/http"
)

var HeaderContent = "Content-Type"
var HeaderKey = "Ocp-Apim-Subscription-Key"
var HeaderLanguage = "Accept-Language"
var HeaderContentValue = "application/json"
var HeaderKeyValue = "befa6b6ab5f4416fbd991b02f5543160"
var HeaderLanguageValue = "en"

func GetCampaignMissions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(HeaderContent, HeaderContentValue)
	// w.Header().Set(HeaderKey, HeaderKeyValue)
	// w.Header().Set(HeaderLanguage, HeaderLanguageValue)
	// var results []primitive.M
	// dbMongo := mongo.Database
	// //slice for multiple documents
	// cur, err := dbMongo.Collection("mission").Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	// if err != nil {

	// 	fmt.Println(err)

	// }
	// for cur.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

	// 	var elem primitive.M
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	results = append(results, elem) // appending document pointed by Next()
	// }
	// cur.Close(context.TODO()) // close the cursor once stream of documents has exhausted
	// json.NewEncoder(w).Encode(results)
}
