package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Celsius int8

type WeatherType string

const (
	THUNDERSTORM WeatherType = "Гроза"
	DRIZZLE      WeatherType = "Морось"
	RAIN         WeatherType = "Дождь"
	SNOW         WeatherType = "Сніг"
	CLEAR        WeatherType = "Ясно"
	FOG          WeatherType = "Туман"
	CLOUDS       WeatherType = "Хмарно"
)

type Weather struct {
	temperature  Celsius
	weather_type WeatherType
	sunrise      time.Time
	sunset       time.Time
	city         string
}

type WeatherData struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Name string `json:"name"`
}

func get_openweather_response(c *Coordinates) string {
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&appid=%s&lang=ru&units=metric",
		c.latitude,
		c.longitude,
		"7549b3ff11a7b2f3cd25b56d21c83c6a",
	)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ApiServiceError:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ApiServiceError:", err)
	}

	return string(body)
}

func parse_openweather_response(r string) Weather {
	var data WeatherData
	err := json.Unmarshal([]byte(r), &data)
	if err != nil {
		fmt.Println("ApiServiceError:", err)
	}

	return Weather{
		temperature:  Celsius(data.Main.Temp),
		weather_type: parse_weather_type(&data),
		sunrise:      time.Unix(data.Sys.Sunrise, 0),
		sunset:       time.Unix(data.Sys.Sunset, 0),
		city:         data.Name}
}

func parse_weather_type(d *WeatherData) WeatherType {
	weather_types := []struct {
		ID   string
		Type WeatherType
	}{
		{"1", THUNDERSTORM},
		{"3", DRIZZLE},
		{"5", RAIN},
		{"6", SNOW},
		{"7", FOG},
		{"800", CLEAR},
		{"80", CLOUDS},
	}

	for _, w := range weather_types {
		if strings.HasPrefix(fmt.Sprint(d.Weather[0].ID), w.ID) {
			return w.Type
		}
	}
	panic("Unknown weather type")
}

func Get_weather(c *Coordinates) Weather {
	response := get_openweather_response(c)
	return parse_openweather_response(response)
}
