package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type WeatherStorage interface {
	save(w *Weather)
}

type PlainFileWeatherStorage struct {
	file string
}

func (p *PlainFileWeatherStorage) save(weather *Weather) {
	now := time.Now().Format("2006-01-02 15:04:05")
	formatted_weather := Format_weather(weather)

	file, err := os.OpenFile(p.file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%s\n%s\n", now, formatted_weather))
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

type JSONFileWeatherStorage struct {
	file string
}

type HistoryRecord struct {
	Date    string `json:"date"`
	Weather string `json:"weather"`
}

func (j *JSONFileWeatherStorage) save(weather *Weather) {
	records := j.read_history()
	now := time.Now().Format("2006-01-02 15:04:05")
	newRecord := HistoryRecord{
		Date:    now,
		Weather: Format_weather(weather),
	}
	records = append(records, newRecord)
	j.write_history(&records)
}

func (j *JSONFileWeatherStorage) read_history() []HistoryRecord {
	file, err := os.OpenFile(j.file, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	var records []HistoryRecord
	err = json.NewDecoder(file).Decode(&records)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	return records
}

func (j *JSONFileWeatherStorage) write_history(records *[]HistoryRecord) {
	file, err := os.OpenFile(j.file, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	file.Seek(0, 0)
	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(records)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func Save_weather(w *Weather, ws WeatherStorage) {
	ws.save(w)
}
