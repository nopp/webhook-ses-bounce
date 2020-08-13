package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type config struct {
	TableName string `json:"tableName"`
	AwsRegion string `json:"awsRegion"`
	HostPort  string `json:"hostPort"`
}

// LoadConfiguration - Read config from config.json
func LoadConfiguration() config {

	var config config

	configFile, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(configFile, &config)
	return config
}

// Message -  Common message to log and mux
func Message(w http.ResponseWriter, message string) {
	log.Println(message)
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}
