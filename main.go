package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/zsombix/StationsGoProject/stations"
)

func main() {

	http.HandleFunc("/stations", getStations)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func getStations(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(&stations.Stations)
	if err != nil {
		log.Fatal("Could not convert")
	}
	//Add the proper header values
	header := w.Header()
	header.Add("Content-Type", "application/json")

	//Write response
	fmt.Fprintf(w, string(data))
}

//This is just an alternative for Encoding to JSON
func tempGetStations(w http.ResponseWriter, r *http.Request) {
	var buffer bytes.Buffer
	json.NewEncoder(&buffer).Encode(&stations.Stations)

	fmt.Fprintf(w, buffer.String())
}
