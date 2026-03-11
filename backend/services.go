package backend

import (
	"encoding/json"
	"net/http"
)

// GetArtists fetches all artists from the API
func GetArtists() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}