package main

import (
	"log"
	"net/http"
	"webhook-ses-bounce/bounce"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// User methods
	router.HandleFunc("/newbounce", bounce.NewBounce).Methods("POST")

	log.Fatal(http.ListenAndServe("0.0.0.0:8181", router))

}
