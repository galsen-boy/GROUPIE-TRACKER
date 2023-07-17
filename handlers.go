package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type PageData struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    []string
	ConcertDates []string
	Relations    map[string][]string
}

type ArtistFullData struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    []string            `json:"locations"`
	ConcertDates []string            `json:"concertDates"`
	Relations    map[string][]string `json:"relations"`
}

type MyArtist struct {
	ID           int                 `json:"id"`
	Image        string              `json:"image"`
	Name         string              `json:"name"`
	Members      []string            `json:"members"`
	CreationDate int                 `json:"creationDate"`
	FirstAlbum   string              `json:"firstAlbum"`
	Locations    string              `json:"locations"`
	ConcertDates string              `json:"concertDates"`
	Relations    map[string][]string `json:"relations"`
}

type MyLocation struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
type LocationData struct {
	Index []MyLocation `json:"index"`
}

type MyConcertDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type ConcertDateData struct {
	Index []MyConcertDate `json:"index"`
}

type MyRelationDate struct {
	ID            int                 `json:"id"`
	DatesLocation map[string][]string `json:"datesLocations"`
}
type RelationData struct {
	Index []MyRelationDate `json:"index"`
}

var Artists []MyArtist
var ArtistData []ArtistFullData
var LocationsData LocationData
var ConcertDatesData ConcertDateData
var RelationsData RelationData

func GetArtistsData() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	if err != nil {
		return errors.New("error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by reading")
	}
	json.Unmarshal(bytes, &Artists)
	return nil
}

func GetLocations() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return errors.New("error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by reading")
	}
	json.Unmarshal(bytes, &LocationsData)
	return nil
}

func GetDates() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return errors.New("error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by reading")
	}
	json.Unmarshal(bytes, &ConcertDatesData)
	return nil
}

func GetRelations() error {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return errors.New("error by get")
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("error by reading")
	}
	json.Unmarshal(bytes, &RelationsData)
	return nil
}

func GetData() {
	GetArtistsData()
	GetLocations()
	GetDates()
	GetRelations()
	var template ArtistFullData
	// var locate MyLocation
	for i := range Artists {
		template.ID = i + 1
		template.Image = Artists[i].Image
		template.Name = Artists[i].Name
		template.Members = Artists[i].Members
		template.CreationDate = Artists[i].CreationDate
		template.FirstAlbum = Artists[i].FirstAlbum
		template.Locations = LocationsData.Index[i].Locations
		template.ConcertDates = ConcertDatesData.Index[i].Dates
		template.Relations = RelationsData.Index[i].DatesLocation
		ArtistData = append(ArtistData, template)
	}

}

func MainHandler(res http.ResponseWriter, req *http.Request) {
	search := req.FormValue("search")
	// if req.URL.Path != "/" || req.Method != "GET" {
	// 	error404(res)
	// 	return
	// }

	if search != "" && len(ArtistData) != 0 {
		ArtistData = Search(search)
	}

	template, _ := template.ParseFiles("./templates/index.html")

	template.Execute(res, ArtistData)
}
func artistHandler(res http.ResponseWriter, req *http.Request) {
	// if req.URL.Path != "/" || req.Method != "GET" {
	// 	error404(res)
	// 	return
	// }
	template, _ := template.ParseFiles("./templates/ArtistPage.html")

	queryParams := req.URL.Query() // Obtient les paramètres de la requête dans un map

	id := queryParams.Get("id")
	data := PageData{}

	for _, v := range ArtistData {
		a, _ := strconv.Atoi(id)
		if v.ID == a {
			data.Name = v.Name
			data.Members = v.Members
			data.Image = v.Image
			data.CreationDate = v.CreationDate
			data.ConcertDates = v.ConcertDates
			data.FirstAlbum = v.FirstAlbum
			data.Locations = v.Locations
			data.Relations = v.Relations
		}
	}

	template.Execute(res, data)
}

func getDatabyId(id int) ArtistFullData {
	var data ArtistFullData

	for i := range Artists {
		if i == id {
			data.ID = Artists[i].ID
			data.Image = Artists[i].Image
			data.Name = Artists[i].Name
			data.Members = Artists[i].Members
			data.CreationDate = Artists[i].CreationDate
			data.FirstAlbum = Artists[i].FirstAlbum
			data.Locations = LocationsData.Index[i].Locations
			data.ConcertDates = ConcertDatesData.Index[i].Dates
			data.Relations = RelationsData.Index[i].DatesLocation
			break
		}
	}
	return data
}

func Search(search string) []ArtistFullData {
	if search == "" {
		return ArtistData
	}
	var resultSearch []ArtistFullData
	search = strings.ToLower(search)
	reg := regexp.MustCompile(`^` + search)
	for i := range Artists {
		temp := strings.ToLower(Artists[i].Name)
		if reg.Match([]byte(temp)) {
			// fmt.Println("search id = ", i)
			resultSearch = append(resultSearch, getDatabyId(i))
		}
	}

	return resultSearch
}
