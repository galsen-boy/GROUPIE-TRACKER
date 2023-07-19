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

	log.Println("http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
