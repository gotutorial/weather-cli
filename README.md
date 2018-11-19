# weather-cli

## Clone Weather CLI

Go to $GOPATH/src and create the following chain of directories

```
$ mkdir $GOPATH/src/github.com/gotutorial
$ cd $GOPATH/src/github.com/gotutorial
```

Clone weather-cli repo in **gotutorial** folder
```
$ git clone https://github.com/gotutorial/weather-cli.git
```

## Build Weather CLI
Use **go build** command build Weather CLI Tool
```
$ cd weather-cli
$ go build -o weather
```

## Run Weather CLI
Use **-h** or **--help** you can see Weather CLI usage help

```
$ ./weather -h
NAME:
   Weather CLI Tool - Prompt the current weather condition for a city

USAGE:
   weather [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --all, -a                  Print all weather attributes
   --city value, -c value     The city name you are looking for its current weather condition
   --country value, -o value  The country name you are looking for its current weather condition
   --desc, -d                 Current description
   --hum, -m                  Current humidity
   --temp, -t                 Current temperature
   --help, -h                 show help
   --version, -v              print the version
```

Use **-c** or **--city** argument you can pass the city name and this command line tool print out the current weather status description in JSON format.

```
$ ./weather -c "New York"
{"location":"New York","description":"overcast clouds"}
```

To see all weather attributes including description, temperature and humidity, use **-a** or **--all**

```
$ ./weather -c "New York"  -a
{"location":"New York","description":"overcast clouds","temperature":39.66,"humidity":100}
```
