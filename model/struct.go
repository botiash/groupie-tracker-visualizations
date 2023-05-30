package model

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int32    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Location struct {
	Id        uint64   `json:"id"`
	Locations []string `json:"locations"`
	DatesUrl  string   `json:"dates"`
}

type Date struct {
	Id    uint64   `json:"id"`
	Dates []string `json:"dates"`
}


