package utils

import (
	"fmt"
	"groupietracker/model"
	"strconv"
	"strings"
)

func Search(text, searchType string) []model.Artist {
	var artists []model.Artist

	switch searchType {
	case "artist":
		artists := SearchArtistByName(text)
		return artists
	case "member":
		artists := SearchArtistByMember(text)
		return artists
	case "location":
		artists := SearchArtistByLocation(text)
		return artists
	case "firstalbum":
		artists := SearchArtistByFirstalbum(text)
		return artists
	case "creationDate":
		artists := SearchArtistByCreationDate(text)
		return artists
	}

	return artists

}

func SearchArtistByFirstalbum(text string) []model.Artist {
	var result []model.Artist
	artists, err := GetArtists()
	if err != nil {
		fmt.Println(err)
		return result
	}
	text = strings.ToLower(text)
	for _, artist := range artists {
		name := strings.ToLower(artist.FirstAlbum)
		if strings.Contains(name, text) {
			result = append(result, artist)
		}
	}
	return result
}

func SearchArtistByCreationDate(text string) []model.Artist {
	var result []model.Artist
	artists, err := GetArtists()
	if err != nil {
		fmt.Println(err)
		return result
	}
	text = strings.ToLower(text)
	for _, artist := range artists {
		date := artist.CreationDate
		searchDate, _ := strconv.Atoi(text)
		if date == searchDate {
			result = append(result, artist)
		}
	}
	return result
}

func SearchArtistByName(text string) []model.Artist {
	var result []model.Artist
	artists, err := GetArtists()
	if err != nil {
		fmt.Println(err)
		return result
	}

	for _, artist := range artists {
		if strings.Contains(artist.Name, text) {
			result = append(result, artist)
		}
	}

	return result

}

func SearchArtistByMember(text string) []model.Artist {
	var result []model.Artist
	artists, err := GetArtists()
	if err != nil {
		fmt.Println(err)
		return result
	}
	for i, artist := range artists {
		for _, member := range artist.Members {
			if strings.Contains(member, text) {
				if !CheckUnique(result, artists[i]) {
					result = append(result, artist)
				}
			}
		}
	}

	return result
}

func SearchArtistByLocation(text string) []model.Artist {
	var result []model.Artist
	artists, err := GetArtists()
	locations, err := GetLocations()
	if err != nil {
		fmt.Println(err)
		return result
	}
	text = strings.ToLower(text)
	for i, location := range locations.Index {
		for _, name := range location.Locations {
			if strings.Contains(name, text) {
				if !CheckUnique(result, artists[i]) {
					result = append(result, artists[i])
				}
			}
		}
	}

	return result
}

func CheckUnique(artists []model.Artist, target model.Artist) bool {
	for _, artist := range artists {
		if artist.Id == target.Id {
			return true
		}
	}
	return false
}
