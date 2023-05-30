package model

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetAllArtists() ([]Artist, int) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil || res.StatusCode != http.StatusOK {
		return []Artist{}, http.StatusInternalServerError
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	// fmt.Println(string(body))
	var Artists []Artist
	err = json.Unmarshal(body, &Artists)
	if err != nil {
		return []Artist{}, http.StatusInternalServerError
	}
	return Artists, http.StatusOK
}

func GetOneArtist(n string) (Artist, int) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + n)
	if err != nil || res.StatusCode != http.StatusOK {
		return Artist{}, http.StatusInternalServerError
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Artist{}, http.StatusInternalServerError
	}
	var OneArtist Artist
	json.Unmarshal(body, &OneArtist)
	return OneArtist, http.StatusOK
}

func GetLocation(n string) (Location, int) {
	var Locations Location

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + n)
	if err != nil || res.StatusCode != http.StatusOK {
		return Location{}, http.StatusInternalServerError
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, http.StatusInternalServerError
	}
	json.Unmarshal(body, &Locations)
	return Locations, http.StatusOK
}

func GetDate(n string) (Date, int) {
	var Dates Date

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + n)
	if err != nil || res.StatusCode != http.StatusOK {
		return Date{}, http.StatusInternalServerError
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Date{}, http.StatusInternalServerError
	}
	json.Unmarshal(body, &Dates)
	return Dates, http.StatusOK
}

func GetRelation(n string) (Relation, int) {
	var Relations Relation

	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + n)
	if err != nil || res.StatusCode != http.StatusOK {
		return Relation{}, http.StatusInternalServerError
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Relation{}, http.StatusInternalServerError
	}
	json.Unmarshal(body, &Relations)
	return Relations, http.StatusOK
}
