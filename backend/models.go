package backend

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
	ID        int      `json:"id"`
	Cities    []string `json:"cities"`
	Countries []string `json:"countries"`
}


type Relation struct {
	ArtistID   int   `json:"artistId"`
	LocationID int   `json:"locationId"`
	Dates      []int `json:"dates"` 
}