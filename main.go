package main

import (
	"log"
	"net/http"
	"webhook-ses-bounce/bounce"
	"webhook-ses-bounce/common"

	"github.com/gorilla/mux"
)

func main() {
	config := common.LoadConfiguration()
	router := mux.NewRouter()
	router.HandleFunc("/newbounce", bounce.PutBounce).Methods("POST")
	log.Fatal(http.ListenAndServe(config.HostPort, router))
}
