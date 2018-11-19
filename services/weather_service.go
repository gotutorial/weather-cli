package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
)

type CurrentWeather struct {
	Location    string  `json:"location"`
	Description string  `json:"description,omitempty"`
	Temprature  float64 `json:"tempurature,omitempty"`
	Humidity    float64 `json:"humidity,omitempty"`
}

func CurrentWeatherByCity(city string, country string, showDescription bool, showTemprature bool, showHumidity bool) (string, error) {
	var url strings.Builder
	url.WriteString("http://api.openweathermap.org/data/2.5/weather")
	url.WriteString("?q=")
	url.WriteString(city)
	if len(country) > 0 {
		url.WriteString(",")
		url.WriteString(country)
	}
	url.WriteString("&units=imperial")
	url.WriteString("&appid=")
	url.WriteString("6bee18b26ef8ab0fb51dfa2b55d79ed5")

	response, err := http.Get(url.String())

	if err != nil {
		log.Fatal(err)
		return "", err
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		jsonParsed, err := gabs.ParseJSON([]byte(string(data)))

		if err != nil {
			log.Fatal(err)
			return "", err
		}

		location, _ := jsonParsed.Path("name").Data().(string)

		var temprature float64
		if showTemprature {
			temprature, _ = jsonParsed.Path("main.temp").Data().(float64)
		}

		var humidity float64
		if showHumidity {
			humidity, _ = jsonParsed.Path("main.humidity").Data().(float64)
		}

		var descriptionString string

		if showDescription {
			descriptions, _ := jsonParsed.Path("weather.description").Children()
			for i, description := range descriptions {
				descriptionString += description.Data().(string)
				if i < len(descriptions)-1 {
					descriptionString += ", "
				}
			}
		}

		currentWeather := &CurrentWeather{
			Location:    location,
			Description: descriptionString,
			Temprature:  temprature,
			Humidity:    humidity,
		}

		currentWeatherJson, err := json.Marshal(currentWeather)
		if err != nil {
			log.Fatal(err)
			return "", err
		}
		return string(currentWeatherJson), nil
	}
}
