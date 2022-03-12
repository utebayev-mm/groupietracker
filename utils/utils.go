package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"groupietracker/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	artistApiUrl   string = "https://groupietrackers.herokuapp.com/api/artists"
	relationApiUrl string = "https://groupietrackers.herokuapp.com/api/relation"
	locationApiUrl string = "https://groupietrackers.herokuapp.com/api/locations"
)

// return Artists from API and UnMarshal json response
func GetArtists() ([]model.Artist, error) {
	req, err := http.Get(artistApiUrl)
	var artists []model.Artist
	if err != nil {
		return artists, err
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return artists, err
	}
	err = json.Unmarshal(body, &artists)
	if err != nil {
		return artists, errors.New("json error")
	}
	return artists, nil
}

func GetLocations() (model.Locations, error) {
	req, err := http.Get(locationApiUrl)
	var locations model.Locations
	if err != nil {
		return locations, err
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return locations, err
	}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return locations, errors.New("json error")
	}
	return locations, nil

}

func GetRelationInfo(id int) (model.Relation, error) {
	var relation model.Relation
	strId := strconv.Itoa(id)
	url := fmt.Sprintf("%v/%v", relationApiUrl, strId)
	req, err := http.Get(url)
	if err != nil {
		return relation, err
	}
	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		return relation, readErr
	}
	err = json.Unmarshal(body, &relation)
	if err != nil {
		return relation, errors.New("json error")
	}
	return relation, nil
}

func GetRelations() (model.Relations, error) {
	var relations model.Relations
	req, err := http.Get(relationApiUrl)
	if err != nil {
		return relations, err
	}
	body, readErr := ioutil.ReadAll(req.Body)
	if readErr != nil {
		return relations, readErr
	}
	err = json.Unmarshal(body, &relations)
	if err != nil {
		return relations, err
	}
	return relations, nil
}

func GetArtistInfo(id int) (model.Artist, error) {
	var artist model.Artist
	strId := strconv.Itoa(id)
	url := fmt.Sprintf("%v/%v", artistApiUrl, strId)
	req, err := http.Get(url)
	if err != nil {
		return artist, err
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return artist, err
	}
	err = json.Unmarshal(body, &artist)
	if err != nil {
		return artist, errors.New("json error")
	}
	if artist.Id == 0 {
		return artist, errors.New("artist not exist")
	}
	return artist, nil
}
