package main

import (
	"log"
	"net/http"
)

func main() {
	GetData()
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/artist/", ArtistHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("http://localhost:8000")

	http.ListenAndServe(":8000", nil)
}
