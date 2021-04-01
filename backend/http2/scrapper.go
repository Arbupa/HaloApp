package http2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/arbupa/HaloApp/structs"
)

func GetMissions() []structs.Missions {

	var MissionStructs []structs.Missions

	req, err := http.NewRequest("GET", "https://www.haloapi.com/metadata/h5/metadata/campaign-missions", nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", "befa6b6ab5f4416fbd991b02f5543160")
	req.Header.Set("Accept-Language", "en")
	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("The HTTP body failed to read with error %s\n", err)
	}

	json.Unmarshal(body, &MissionStructs)
	return MissionStructs
}

func GetMaps() []structs.Maps {

	var MapStructs []structs.Maps

	req, err := http.NewRequest("GET", "https://www.haloapi.com/metadata/h5/metadata/maps", nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", "befa6b6ab5f4416fbd991b02f5543160")
	req.Header.Set("Accept-Language", "en")
	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("The HTTP body failed to read with error %s\n", err)
	}

	json.Unmarshal(body, &MapStructs)

	return MapStructs
}

func GetSkulls() []structs.Skulls {

	var SkullStructs []structs.Skulls

	req, err := http.NewRequest("GET", "https://www.haloapi.com/metadata/h5/metadata/skulls", nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", "befa6b6ab5f4416fbd991b02f5543160")
	req.Header.Set("Accept-Language", "en")
	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("The HTTP body failed to read with error %s\n", err)
	}

	json.Unmarshal(body, &SkullStructs)

	return SkullStructs
}

func GetEnemies() []structs.Enemies {
	var EnemiesStructs []structs.Enemies

	req, err := http.NewRequest("GET", "https://www.haloapi.com/metadata/h5/metadata/enemies", nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", "befa6b6ab5f4416fbd991b02f5543160")
	req.Header.Set("Accept-Language", "en")
	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("The HTTP body failed to read with error %s\n", err)
	}

	json.Unmarshal(body, &EnemiesStructs)

	return EnemiesStructs
}

func GetWeapons() []structs.Weapons {
	var WeaponsStructs []structs.Weapons

	req, err := http.NewRequest("GET", "https://www.haloapi.com/metadata/h5/metadata/weapons", nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", "befa6b6ab5f4416fbd991b02f5543160")
	req.Header.Set("Accept-Language", "en")
	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("The HTTP body failed to read with error %s\n", err)
	}

	json.Unmarshal(body, &WeaponsStructs)

	return WeaponsStructs
}

func GetVehicles() []structs.Vehicles {

	var vehicles []structs.Vehicles

	req, err := http.NewRequest("GET", "https://www.haloapi.com/metadata/h5/metadata/vehicles", nil)
	req.Header.Set("Ocp-Apim-Subscription-Key", "befa6b6ab5f4416fbd991b02f5543160")
	req.Header.Set("Accept-Language", "en")
	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("The HTTP body failed to read with error %s\n", err)
	}

	json.Unmarshal(body, &vehicles)

	return vehicles
}
