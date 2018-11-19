package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/gotutorial/weather-cli/services"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp() // := is the Short variable declarations
	app.Name = "Weather CLI Tool"
	app.Usage = "Prompt the current weather condition for a city"
	app.Version = "1.0.0"
	app.HideHelp = false

	var city string
	var country string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "city, c",
			Usage:       "The city name you are looking for its current weather condition",
			Destination: &city,
		},
		cli.StringFlag{
			Name:        "country, o",
			Usage:       "The country name you are looking for its current weather condition",
			Destination: &country,
		},
		cli.BoolFlag{
			Name:  "temp, t",
			Usage: "Current temperature",
		},
		cli.BoolFlag{
			Name:  "desc, d",
			Usage: "Current description",
		},
		cli.BoolFlag{
			Name:  "hum, m",
			Usage: "Current humidity",
		},
		cli.BoolFlag{
			Name:  "all, a",
			Usage: "Print all weather attributes",
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(city) > 0 {
			var showDescription bool
			showDescription = false
			if !(c.Bool("desc") || c.Bool("temp") || c.Bool("hum")) || c.Bool("desc") || c.Bool("all") {
				showDescription = true
			}
			var reponse, err = services.CurrentWeatherByCity(city, country, showDescription, c.Bool("temp") || c.Bool("all"), c.Bool("hum") || c.Bool("all"))
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Println(reponse)
			}
		} else {
			fmt.Println("City is required. Pass the city value using -c or --city")
			cli.ShowAppHelp(c)
		}

		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
