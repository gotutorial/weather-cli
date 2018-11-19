package services

import (
	"testing"

	"github.com/nbio/st"
	gock "gopkg.in/h2non/gock.v1"
)

func TestShowAllWeatherAttributes(t *testing.T) {
	defer gock.Off()

	gock.New("http://api.openweathermap.org").
		Get("/data/2.5/weather").
		Reply(200).
		File("weather_api_response_200.json")

	currentWeatherJson, _ := CurrentWeatherByCity("Atlanta", "", true, true, true)

	st.Expect(t, currentWeatherJson, `{"location":"Columbus","description":"clear sky","tempurature":50.73,"humidity":81}`)

	// Verify that we don't have pending mocks
	st.Expect(t, gock.IsDone(), true)
}

func TestShowDescription(t *testing.T) {
	defer gock.Off()

	gock.New("http://api.openweathermap.org").
		Get("/data/2.5/weather").
		Reply(200).
		File("weather_api_response_200.json")

	currentWeatherJson, _ := CurrentWeatherByCity("Atlanta", "", true, false, false)

	st.Expect(t, currentWeatherJson, `{"location":"Columbus","description":"clear sky"}`)

	// Verify that we don't have pending mocks
	st.Expect(t, gock.IsDone(), true)
}

func TestShowTemprature(t *testing.T) {
	defer gock.Off()

	gock.New("http://api.openweathermap.org").
		Get("/data/2.5/weather").
		Reply(200).
		File("weather_api_response_200.json")

	currentWeatherJson, _ := CurrentWeatherByCity("Atlanta", "", false, true, false)

	st.Expect(t, currentWeatherJson, `{"location":"Columbus","tempurature":50.73}`)

	// Verify that we don't have pending mocks
	st.Expect(t, gock.IsDone(), true)
}

func TestShowHumidity(t *testing.T) {
	defer gock.Off()

	gock.New("http://api.openweathermap.org").
		Get("/data/2.5/weather").
		Reply(200).
		File("weather_api_response_200.json")

	currentWeatherJson, _ := CurrentWeatherByCity("Atlanta", "", false, false, true)

	st.Expect(t, currentWeatherJson, `{"location":"Columbus","humidity":81}`)

	// Verify that we don't have pending mocks
	st.Expect(t, gock.IsDone(), true)
}
