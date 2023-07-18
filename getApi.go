package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

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
