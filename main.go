package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	}
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		}
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text         string  `json:"text"`
					ChanceOfRain float64 `json:"chance_of_rain"`
				}
			}
		}
	} `json:"forecast"`
}

func main() {
	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=d8dab35c10aa4b57b85160946242405&q=Omsk&days=1")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	location, current, _ := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour[0]
	fmt.Printf("%s, %s: %.0fC, %s\n", location.Name, location.Country, current.TempC, current.Condition.Text)
}
