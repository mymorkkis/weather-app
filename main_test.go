package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testServer(mockJSONResponse string) *httptest.Server {
	th := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(mockJSONResponse))
	})
	ts := httptest.NewServer(th)
	return ts
}

func TestFetchData(t *testing.T) {
	ts := testServer(mockJSONResponse)
	defer ts.Close()

	weatherResponseData := fetchCityWeatherData(ts.URL)

	if !cmp.Equal(expectedWeatherAPIData, weatherResponseData) {
		t.Error("Expected weather JSON response not returned")
	}
}

var expectedWeatherAPIData = cityWeatherAPIResponse{
	CityName:    "Lisbon",
	Lon:         "-9.13333",
	Timezone:    "Europe/Lisbon",
	Lat:         "38.71667",
	CountryCode: "PT",
	Data: []weatherReport{
		{
			TimestampUTC:   "2020-03-15T08:00:00",
			TimestampLocal: "2020-03-15T08:00:00",
			UV:             1.54664,
			Weather: weather{
				Description: "Few clouds",
			},
		},
		{
			TimestampUTC:   "2020-03-15T09:00:00",
			TimestampLocal: "2020-03-15T09:00:00",
			UV:             2.22495,
			Weather: weather{
				Description: "Few clouds",
			},
		},
	},
}

const mockJSONResponse = `{
    "data": [
        {
            "wind_cdir": "NNW",
            "rh": 82,
            "pod": "d",
            "timestamp_utc": "2020-03-15T08:00:00",
            "pres": 1016.75,
            "solar_rad": 174.133,
            "ozone": 346.2,
            "weather": {
                "icon": "c02d",
                "code": 801,
                "description": "Few clouds"
            },
            "wind_gust_spd": 7.01649,
            "timestamp_local": "2020-03-15T08:00:00",
            "snow_depth": 0,
            "clouds": 43,
            "ts": 1584259200,
            "wind_spd": 4.18348,
            "pop": 0,
            "wind_cdir_full": "north-northwest",
            "slp": 1018,
            "dni": 559.79,
            "dewpt": 9.5,
            "snow": 0,
            "uv": 1.54664,
            "wind_dir": 350,
            "clouds_hi": 0,
            "precip": 0,
            "vis": 24.1349,
            "dhi": 62.67,
            "app_temp": 12.2,
            "datetime": "2020-03-15:08",
            "temp": 12.2,
            "ghi": 182.52,
            "clouds_mid": 0,
            "clouds_low": 43
        },
        {
            "wind_cdir": "NNW",
            "rh": 77,
            "pod": "d",
            "timestamp_utc": "2020-03-15T09:00:00",
            "pres": 1017.24,
            "solar_rad": 375.074,
            "ozone": 345.323,
            "weather": {
                "icon": "c02d",
                "code": 801,
                "description": "Few clouds"
            },
            "wind_gust_spd": 6.94322,
            "timestamp_local": "2020-03-15T09:00:00",
            "snow_depth": 0,
            "clouds": 32,
            "ts": 1584262800,
            "wind_spd": 4.26229,
            "pop": 0,
            "wind_cdir_full": "north-northwest",
            "slp": 1018.51,
            "dni": 733.28,
            "dewpt": 9.2,
            "snow": 0,
            "uv": 2.22495,
            "wind_dir": 352,
            "clouds_hi": 0,
            "precip": 0,
            "vis": 24.1351,
            "dhi": 86.13,
            "app_temp": 13.2,
            "datetime": "2020-03-15:09",
            "temp": 13.2,
            "ghi": 381.01,
            "clouds_mid": 0,
            "clouds_low": 32
        }

    ],
    "city_name": "Lisbon",
    "lon": "-9.13333",
    "timezone": "Europe/Lisbon",
    "lat": "38.71667",
    "country_code": "PT",
    "state_code": "14"
}`
