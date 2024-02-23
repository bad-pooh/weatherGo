package main

type Coordinates struct {
	latitude  float64
	longitude float64
}

func GetGpsCoordinates() Coordinates {
	return Coordinates{49.2, 28.4}
}
