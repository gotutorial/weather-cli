package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func weatherByCity() {
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=Columbus&appid=6bee18b26ef8ab0fb51dfa2b55d79ed5")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
