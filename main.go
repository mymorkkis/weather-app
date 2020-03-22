package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymorkkis/weather-api/utils"
)

const weatherbitBaseURL = "https://api.weatherbit.io/v2.0/forecast/hourly?"

func main() {
	args := parseCommandLineArguments()
	args.apiKey = getAPIKey()

	requestParamaters := fmt.Sprintf(
		"city=%s&key=%s&hours=%d&lang=%s",
		args.city, args.apiKey, args.hours, args.language,
	)
	url := fmt.Sprintf("%s%s", weatherbitBaseURL, requestParamaters)

	cityWeatherData := fetchCityWeatherData(url)

	log.Print("cityWeatherData: ", cityWeatherData)
}

func parseCommandLineArguments() arguments {
	city := flag.String("city", "lisbon", "City to provide weather forecast for")
	hours := flag.Int("hours", 2, "Hours of forcast to return (maximum 48)")
	language := flag.String("lan", "en", "Two letter language code")

	flag.Parse()

	return arguments{
		city:     *city,
		hours:    *hours,
		language: *language,
	}
}

func getAPIKey() (apiKey string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey = os.Getenv("WEATHERBIT_API_KEY")

	if apiKey == "" {
		log.Fatal("WEATHERBIT_API_KEY not set in .env file")
	}

	return
}

func fetchCityWeatherData(url string) (weather cityWeatherAPIResponse) {
	weather = cityWeatherAPIResponse{}
	data := utils.FetchData(url)
	utils.ParseFromJSON(data, &weather)
	return
}

type arguments struct {
	apiKey   string
	city     string
	language string
	hours    int
}

type cityWeatherAPIResponse struct {
	Data        []weatherReport
	CityName    string `json:"city_name"`
	Timezone    string
	Lat         string
	Lon         string
	CountryCode string `json:"country_code"`
}

type weatherReport struct {
	TimestampUTC   string `json:"timestamp_utc"`
	TimestampLocal string `json:"timestamp_local"`
	UV             float32
	Weather        weather
}

type weather struct {
	Description string
}
