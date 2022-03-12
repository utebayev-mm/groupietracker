package controller

import (
	"errors"
	"fmt"
	"groupietracker/model"
	"groupietracker/utils"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// const emptyQueryString string = "creationdateFrom=&creationdateTo=&firstalbumFrom=&firstalbumTo=&numbersofmemberFrom=&numbersofmemberTo=&location=&submit="

// MainPage - Controller for handle main page requests
func MainPage(w http.ResponseWriter, r *http.Request) {
	// handle only get and route /
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodPost {
		r.ParseForm()
		search_text := r.FormValue("search")
		search_type := r.FormValue("type")
		searched := utils.Search(search_text, search_type)
		artists, err := utils.GetArtists()
		locations, err := utils.GetLocations()
		var MainInfo = model.MainInfo{
			Artists:   artists,
			Locations: locations,
			Searched:  searched,
		}
		// fmt.Println(searched)
		if err != nil {
			fmt.Println(err)
			ErrorHandler(w, http.StatusNotFound)
			return
		}
		tmpl, err := template.ParseFiles("./static/templates/index.html")
		if err != nil {
			fmt.Println(err)
			ErrorHandler(w, http.StatusNotFound)
			return
		}
		if err := tmpl.Execute(w, MainInfo); err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}

	}
	if r.Method == http.MethodGet {
		artists, err := utils.GetArtists()
		locations, err := utils.GetLocations()
		// fmt.Println(r.URL.RawQuery)
		// fmt.Println(len(r.URL.RawQuery))
		if r.URL.RawQuery != "" {
			var err2 error
			// artists = Filter(r)
			artists, err2 = Filter(r, artists, w)
			//	fmt.Println("We are filtered")
			if err2 != nil {
				fmt.Println(err2)
				ErrorHandler(w, http.StatusBadRequest)
				return
			}
		} else {
			fmt.Println("no filter")
		}
		var MainInfo = model.MainInfo{
			Artists:   artists,
			Locations: locations,
			Searched:  artists,
		}
		if err != nil {
			fmt.Println(err)
			ErrorHandler(w, http.StatusNotFound)
			return
		}
		tmpl, err := template.ParseFiles("./static/templates/index.html")
		if err != nil {
			fmt.Println(err)
			ErrorHandler(w, http.StatusNotFound)
			return
		}
		if err := tmpl.Execute(w, MainInfo); err != nil {
			ErrorHandler(w, http.StatusInternalServerError)
			return
		}
	}

}

func Filter(r *http.Request, artists []model.Artist, w http.ResponseWriter) ([]model.Artist, error) {
	var Filtered []model.Artist
	err := FilterByCreationDate(r, artists, &Filtered, w)
	if err != nil {
		return nil, err
	}
	Filtered, err = FilterByFirstAlbumDate(r, Filtered, w)
	if err != nil {
		return nil, err
	}
	Filtered, err = FilterByMembers(r, Filtered, w)
	if err != nil {
		return nil, err
	}
	Filtered = FilterByLocation(r, Filtered)
	return Filtered, err
}

func FilterByLocation(r *http.Request, Filtered []model.Artist) []model.Artist {
	artists := Filtered
	locationText := r.URL.Query().Get("location")
	if locationText == "" {
		return Filtered
	}
	text := strings.ToLower(locationText)
	var arr []int
	relations, _ := utils.GetRelations()
	for i, artist := range artists {
		relInfo := relations.Index[artist.Id-1]
		for city := range relInfo.DatesLocations {
			city = strings.ToLower(strings.Split(city, "-")[0] + ", " + strings.Split(city, "-")[1])
			if strings.Contains(city, text) {
				arr = append(arr, i)
			}
		}
	}
	var Filtered2 []model.Artist
	var arr2 []int
	for i := 0; i < len(arr); i++ {
		counter := 0
		for j := 0; j < len(arr2); j++ {
			if arr[i] == arr2[j] {
				counter++
			}
		}
		if counter == 0 {
			arr2 = append(arr2, arr[i])
		}
	}
	for i := 0; i < len(arr2); i++ {
		Filtered2 = append(Filtered2, Filtered[arr2[i]])
	}
	return Filtered2
}

func FilterByMembers(r *http.Request, Filtered []model.Artist, w http.ResponseWriter) ([]model.Artist, error) {
	a := []rune(r.URL.Query().Get("numbersofmemberFrom"))
	b := []rune(r.URL.Query().Get("numbersofmemberTo"))
	for _, letter := range a {
		if letter >= 48 && letter <= 57 {
			//	fmt.Println("numbersofmember from is ok")
		} else {
			return nil, errors.New("bad request")
		}
	}
	for _, letter := range b {
		if letter >= 48 && letter <= 57 {
			//	fmt.Println("numbersofmember to is ok")
		} else {
			return nil, errors.New("bad request")
		}
	}
	numbersofmemberFrom, _ := strconv.Atoi(r.URL.Query().Get("numbersofmemberFrom"))
	numbersofmemberTo, _ := strconv.Atoi(r.URL.Query().Get("numbersofmemberTo"))
	if numbersofmemberFrom == 0 && numbersofmemberTo == 0 {
		return Filtered, nil
	}
	if numbersofmemberFrom == 0 {
		numbersofmemberFrom++
	}
	if numbersofmemberFrom > numbersofmemberTo {
		numbersofmemberTo = 30
	}
	var arr []int
	for i, artist := range Filtered {
		number := len(artist.Members)
		if number >= numbersofmemberFrom && number <= numbersofmemberTo {
			// fmt.Println(artist.Name)
			arr = append(arr, i)
		} else {
			//arr = append(arr, i)
		}
	}
	var Filtered2 []model.Artist
	for i := 0; i < len(arr); i++ {
		Filtered2 = append(Filtered2, Filtered[arr[i]])
	}
	return Filtered2, nil
}

func FilterByFirstAlbumDate(r *http.Request, Filtered []model.Artist, w http.ResponseWriter) ([]model.Artist, error) {
	a := []rune(r.URL.Query().Get("firstalbumFrom"))
	b := []rune(r.URL.Query().Get("firstalbumTo"))
	for _, letter := range a {
		if letter >= 48 && letter <= 57 {
			//	fmt.Println("first album from is ok")
		} else {
			return nil, errors.New("bad request")
		}
	}
	for _, letter := range b {
		if letter >= 48 && letter <= 57 {
			//	fmt.Println("first album to is ok")
		} else {
			return nil, errors.New("bad request")
		}
	}
	firstalbumFrom, _ := strconv.Atoi(r.URL.Query().Get("firstalbumFrom"))
	firstalbumTo, _ := strconv.Atoi(r.URL.Query().Get("firstalbumTo"))
	if firstalbumFrom == 0 && firstalbumTo == 0 {
		return Filtered, nil
	}
	if firstalbumFrom == 0 {
		firstalbumFrom = 1950
	}
	if firstalbumFrom > firstalbumTo {
		firstalbumTo = 2100
	}
	var arr []int
	for i, artist := range Filtered {
		date, _ := strconv.Atoi(artist.FirstAlbum[len(artist.FirstAlbum)-4:])
		if date >= firstalbumFrom && date <= firstalbumTo {
			arr = append(arr, i)
		} else {

		}
	}
	var Filtered2 []model.Artist
	//fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		Filtered2 = append(Filtered2, Filtered[arr[i]])
	}
	return Filtered2, nil
}

func FilterByCreationDate(r *http.Request, artists []model.Artist, Filtered *[]model.Artist, w http.ResponseWriter) error {
	a := []rune(r.URL.Query().Get("creationdateFrom"))
	b := []rune(r.URL.Query().Get("creationdateTo"))
	for _, letter := range a {
		if letter >= 48 && letter <= 57 {
			//fmt.Println("creation date from is ok")
		} else {
			return errors.New("bad request")
		}
	}
	for _, letter := range b {
		if letter >= 48 && letter <= 57 {
			//fmt.Println("creation date to is ok")
		} else {
			return errors.New("bad request")
		}
	}
	creationdateFrom, _ := strconv.Atoi(r.URL.Query().Get("creationdateFrom"))
	creationdateTo, _ := strconv.Atoi(r.URL.Query().Get("creationdateTo"))
	if creationdateFrom == 0 {
		creationdateFrom = 1950
	}
	if creationdateFrom > creationdateTo {
		creationdateTo = 2100
	}
	for _, artist := range artists {
		date := artist.CreationDate
		if date >= creationdateFrom && date <= creationdateTo {
			*Filtered = append(*Filtered, artist)
		}
	}
	return nil
}

func ErrorHandler(w http.ResponseWriter, Status int) {
	tmpl, err := template.ParseFiles("./static/templates/error.html")
	if err != nil {
		// ErrorHandler(w, http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// fmt.Println("error")
		return
	}
	w.WriteHeader(Status)
	if err := tmpl.Execute(w, http.StatusText(Status)); err != nil {
		// ErrorHandler(w, http.StatusInternalServerError)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		// fmt.Println("error")
		return
	}
}

// ArtistPage - Get more information for specific artist
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	strId := r.URL.Path[8:]
	id, err := strconv.Atoi(strId)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	artistInfo, err := utils.GetArtistInfo(id)
	if err != nil {
		fmt.Println(err)
		if err.Error() == "artist not exist" {
			ErrorHandler(w, http.StatusNotFound)
			fmt.Println("page not found")
			return
		}
		ErrorHandler(w, http.StatusInternalServerError)
		return

	}
	datesLocations, _ := utils.GetRelationInfo(id)

	MoreInfo := model.ArtistInfo{
		Artist:   artistInfo,
		Relation: datesLocations,
	}
	tmpl, err := template.ParseFiles("./static/templates/artist.html")
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	tmpl.Execute(w, MoreInfo)
}
