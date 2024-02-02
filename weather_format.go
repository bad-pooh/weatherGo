package main

import "fmt"

func Format_weather(w *Weather) string {
	return fmt.Sprintf("%s, температура %d°C, %s\nСхід: %s\nЗахід: %s\n",
		w.city,
		w.temperature,
		w.weather_type,
		w.sunrise.Format("15:04"),
		w.sunset.Format("15:04"),
	)
}
