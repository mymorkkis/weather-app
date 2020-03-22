package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	reqHeader    = "weather-api-test"
	timeoutLimit = time.Second * 20
)

// FetchData returns http request response for given URL. Logs Fatal if error.
func FetchData(url string) (res *http.Response) {
	client := http.Client{Timeout: timeoutLimit}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	checkErr(err, "Error creating http request")

	req.Header.Set("User-Agent", reqHeader)

	res, err = client.Do(req)
	checkErr(err, "Error sending http request")

	return
}

// ParseFromJSON parses http response to given struct type. Logs Fatal if error.
func ParseFromJSON(data *http.Response, toStruct interface{}) {
	body, err := ioutil.ReadAll(data.Body)
	checkErr(err, "Error reading http data")

	err = json.Unmarshal(body, &toStruct)
	checkErr(err, "Error unmarshalling json")
}

func checkErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
	}
}
