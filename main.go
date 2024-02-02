package main

import "fmt"

func main() {
	coordinates := Get_gps_coordinates()
	weather := Get_weather(&coordinates)
	fmt.Println(Format_weather(&weather))
	Save_weather(&weather, &PlainFileWeatherStorage{file: "weather_data.txt"})
	Save_weather(&weather, &JSONFileWeatherStorage{file: "weather_data.json"})
}
