package stations

//Station that is used in the project
type Station struct {
	ID        int32
	Name      string
	Latitude  float64
	Longitude float64
}

//Stations is a list of available stations
var Stations [3]Station = [3]Station{
	{
		ID:        1,
		Name:      "Budapest",
		Latitude:  47.497913,
		Longitude: 19.040236,
	},
	{
		ID:        2,
		Name:      "London",
		Latitude:  51.507351,
		Longitude: -0.127758,
	},
	{
		ID:        3,
		Name:      "Bucharest",
		Latitude:  44.426765,
		Longitude: 26.102537,
	},
}
