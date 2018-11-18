package services

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Jeffail/gabs"
)

type CurrentWeather struct {
	location    string
	description string
	tempurature string
	humidity    string
}

func CurrentWeatherByCity(city string, country string, description bool, temprature bool, humidity bool) (*CurrentWeather, error) {
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
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		jsonParsed, err := gabs.ParseJSON([]byte(string(data)))

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		location, _ := jsonParsed.Path("name").Data().(string)

		currentWeather := &CurrentWeather{
			location:    location,
			description: "string",
			tempurature: "string",
			humidity:    "string",
		}
		return currentWeather, nil
	}
}
