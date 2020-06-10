package main

import (
	// "bytes"
	// "encoding/json"
	// "fmt"
	// "log"
	// "net/http"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/zsombix/StationsGoProject/controllers"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	app.Static("/", "./FrontEnd")

	app.Get("/stations", controllers.GetAllStations)

	app.Listen(3000)
}

/*
	This is just what I used before swithing to Fiber
*/

// func main() {

// 	http.HandleFunc("/stations", getStations)
// 	log.Fatal(http.ListenAndServe(":8080", nil))

// }

// func getStations(w http.ResponseWriter, r *http.Request) {
// 	data, err := json.Marshal(&stations.Stations)
// 	if err != nil {
// 		log.Fatal("Could not convert")
// 	}
// 	//Add the proper header values
// 	header := w.Header()
// 	header.Set("Content-Type", "application/json")
// 	header.Set("Access-Control-Allow-Origin", "*")

// 	//Write response
// 	fmt.Fprintf(w, string(data))
// }

// //This is just an alternative for Encoding to JSON
// func tempGetStations(w http.ResponseWriter, r *http.Request) {
// 	var buffer bytes.Buffer
// 	json.NewEncoder(&buffer).Encode(&stations.Stations)

// 	fmt.Fprintf(w, buffer.String())
// }
