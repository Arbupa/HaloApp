package structs

type Maps struct {
	ID                 string
	IDApi              string   `json:"id"`
	Name               string   `json:"name"`
	Description        string   `json:"description"`
	ImageUrl           string   `json:"imageUrl"`
	SupportedGameModes []string `json:"supportedGameModes"`
}

type Skulls struct {
	ID          string
	IDApi       string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageUrl    string `json:"imageUrl"`
	// buscar esta madre
	MissionID string `json:"missionId"`
}

type Missions struct {
	ID            string
	IDApi         string `json:"id"`
	MissionNumber int    `json:"missionNumber"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ImageUrl      string `json:"imageUrl"`
	Type          string `json:"type"`
}

type Enemies struct {
	ID       string
	IDApi    string `json:"id"`
	Faction  string `json:"faction"`
	Name     string `json:"name"`
	ImageUrl string `json:"largeIconImageUrl"`
}

type Weapons struct {
	ID               string
	IDApi            string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ImageUrl         string `json:"largeIconImageUrl"`
	IsUsableByPlayer bool   `json:"isUsableByPlayer"`
	Type             string `json:"type"`
}

type Vehicles struct {
	ID               string
	IDApi            string `json:"id"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ImageUrl         string `json:"largeIconImageUrl"`
	IsUsableByPlayer bool   `json:"isUsableByPlayer"`
}
