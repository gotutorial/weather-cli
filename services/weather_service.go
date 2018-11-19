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
	Location    string `json:"location"`
	Description string `json:"description"`
	Tempurature string `json:"tempurature"`
	Humidity    string `json:"humidity"`
}

func CurrentWeatherByCity(city string, country string, description bool, temprature bool, humidity bool) (string, error) {
	var url strings.Builder
	url.WriteString("http://api.openweathermap.org/data/2.5/weather")
	url.WriteString("?q=")
	url.WriteString(city)
	if len(country) > 0 {
		url.WriteString(",")
		url.WriteString(country)
	}
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
		descriptions, _ := jsonParsed.Path("weather.description").Children()

		var descriptionString string

		for i, description := range descriptions {
			descriptionString += description.Data().(string)
			if i < len(descriptions)-1 {
				descriptionString += ", "
			}
		}

		currentWeather := &CurrentWeather{
			Location:    location,
			Description: descriptionString,
			Tempurature: "string",
			Humidity:    "string",
		}

		currentWeatherJson, err := json.Marshal(currentWeather)
		if err != nil {

			log.Fatal(err)
			return "", err
		}
		return string(currentWeatherJson), nil
	}
}
