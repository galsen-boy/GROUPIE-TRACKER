package main

import (
	"log"
	"net/http"
)

func main() {
	GetData()
	http.HandleFunc("/", MainHandler)
	http.HandleFunc("/artist/", artistHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("http://localhost:9000")

	http.ListenAndServe(":9000", nil)
}
