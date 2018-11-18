package services

import (
	"testing"

	"github.com/nbio/st"
	gock "gopkg.in/h2non/gock.v1"
)

func TestSimple(t *testing.T) {
	defer gock.Off()

	gock.New("http://api.openweathermap.org").
		Get("/data/2.5/weather").
		Reply(200).
		File("weather_api_response_200.json")

	var forcast *WeatherForcast
	forcast, _ = WeatherByCity("Atlanta", "", true, true, true)

	st.Expect(t, forcast.location, `Atlanta`)
	// Verify that we don't have pending mocks
	st.Expect(t, gock.IsDone(), true)
}
