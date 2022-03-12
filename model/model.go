package model

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstalbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Relation struct {
	Id             uint64              `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Relations struct {
	Index []struct {
		Id             uint64              `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type ArtistInfo struct {
	Artist
	Relation
}

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type MainInfo struct {
	Locations Locations
	Artists   []Artist
	Searched  []Artist
}
