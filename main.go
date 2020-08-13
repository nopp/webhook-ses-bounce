package main

import (
	"log"
	"net/http"
	"webhook-ses-bounce/bounce"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/newbounce", bounce.PutBounce).Methods("POST")
	log.Fatal(http.ListenAndServe("0.0.0.0:8181", router))
}
