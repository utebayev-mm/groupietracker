package main

import (
	"fmt"
	"groupietracker/controller"
	"log"
	"net/http"
	"time"
)

func main() {
	// server conf
	Addr := ":8080"
	s := &http.Server{
		Addr:           Addr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// routes
	http.HandleFunc("/artist/", controller.ArtistPage)
	http.HandleFunc("/", controller.MainPage)
	http.Handle("/static/",http.StripPrefix("/static",http.FileServer(http.Dir("./static"))))

	// log
	fmt.Printf("Server listens http://localhost%v\n", Addr)
	log.Fatal(s.ListenAndServe())
}
