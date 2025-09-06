package weather

import (
	"encoding/json"
	"fmt"
	"github.com/maryam-nokohan/WeatherApp/backend/internal/models"
	"io"
	"net/http"
)

func NewCity() *models.City {
	return &models.City{
		CountryName: "none",
		CityName:    "none",
		Latitude:    0,
		Longitude:   0,
	}
}

func GetCity(cityName string) (*models.City, error) {
	fmt.Printf("GetCity: searching for %s\n", cityName)
	url := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1", cityName)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("faild to fetch th coordinates")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("geocoding api error : %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("faild to read the response")
	}

	var geo models.GeoResponse

	if err := json.Unmarshal(body, &geo); err != nil {
		return nil, fmt.Errorf("faild to unmarshal the json api")
	}

	if len(geo.Results) == 0 {
		return nil, fmt.Errorf("city not found: %s", cityName)
	}
	fmt.Printf("GetCity: result: %+v\n", &geo.Results[0])
	return &geo.Results[0], nil

}
