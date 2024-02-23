package main

import "fmt"

func main() {
	coordinates := GetGpsCoordinates()
	weather := GetWeather(&coordinates)
	fmt.Println(FormatWeather(&weather))
	SaveWeather(&weather, &PlainFileWeatherStorage{file: "weather_data.txt"})
	SaveWeather(&weather, &JSONFileWeatherStorage{file: "weather_data.json"})
}
