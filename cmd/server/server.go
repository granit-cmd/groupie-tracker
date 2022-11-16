package server

import (
	"fmt"
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("/ui"))
	mux.Handle("/ui/", http.StripPrefix("/ui", fileServer))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/artist", detailArtist)

	fmt.Println("Starting the web server on http://localhost:4200")
	err := http.ListenAndServe(":4200", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}
