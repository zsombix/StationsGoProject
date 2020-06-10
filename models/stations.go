package stations

import (
	"github.com/Kamva/mgm/v2"
)

//Station that is used in the project
type Station struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string  `json:"name" bson:"name"`
	Latitude         float64 `json:"latitude" bson:"latitude"`
	Longitude        float64 `json:"longitude" bson:"longitude"`
}

//CreateStation : Use this to create a new station object
func CreateStation(name string, latitude, longitude float64) *Station {
	return &Station{
		Name:      name,
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func (model *Station) CollectionName() string {
	return "stations2"
}

//Stations is a list of available stations
var Stations [3]Station = [3]Station{
	{
		Name:      "Budapest",
		Latitude:  47.497913,
		Longitude: 19.040236,
	},
	{
		Name:      "London",
		Latitude:  51.507351,
		Longitude: -0.127758,
	},
	{
		Name:      "Bucharest",
		Latitude:  44.426765,
		Longitude: 26.102537,
	},
}
