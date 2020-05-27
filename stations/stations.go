package stations

//Station that is used in the project
type Station struct {
	Name      string
	Latitude  float64
	Longitude float64
}

//Stations is a list of available stations
var Stations [3]Station = [3]Station{
	Station{
		Name:      "Budapest",
		Latitude:  12.3333,
		Longitude: 43.123,
	},
	Station{
		Name:      "Budapest",
		Latitude:  12.3333,
		Longitude: 43.123,
	},
	Station{
		Name:      "Budapest",
		Latitude:  12.3333,
		Longitude: 43.123,
	},
}
