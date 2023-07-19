package main

import (
	"html/template"
	"net/http"
	"strconv"
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
	if req.URL.Path != "/" {
		error404(res)
		return
	}
	if req.Method != "GET" {
		error405(res)
		return
	}
	template, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		error500(res)
		return
	}
	res.WriteHeader(200)
	template.Execute(res, ArtistData)
}
func ArtistHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/artist/" {
		NotFound(res)
		return
	}
	if req.Method != "GET" {
		error405(res)
		return
	}
	template, err := template.ParseFiles("./templates/ArtistPage.html")
	if err != nil {
		error500(res)
		return
	}
	queryParams := req.URL.Query() // Obtient les paramètres de la requête dans un map
	id := queryParams.Get("id")
	data := PageData{}
	a, err := strconv.Atoi(id)
	for _, v := range ArtistData {
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
	if err != nil || a < 0 {
		error400(res)
		return
	}
	res.WriteHeader(200)
	template.Execute(res, data)
}
