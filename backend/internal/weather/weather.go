package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/maryam-nokohan/WeatherApp/backend/internal/models"
)

type WeatherResponse struct {
	Latitude             float64       `json:"latitude"`
	Longitude            float64       `json:"longitude"`
	Elevation            float64       `json:"elevation"`
	Timezone             string        `json:"timezone"`
	TimezoneAbbreviation string        `json:"timezone_abbreviation"`
	Hourly               models.Hourly `json:"hourly"`
	HourlyUnits          models.Units  `json:"hourly_units"`
}

type Weather interface {
	GetWeather(lat float64, lon float64) (*WeatherResponse, error)
}

func NewWeatherResponse() *WeatherResponse {
	return &WeatherResponse{
		Latitude:             0.0,
		Longitude:            0.0,
		Elevation:            0.0,
		Timezone:             "none",
		TimezoneAbbreviation: "None",
		Hourly: models.Hourly{
			Temperature2m: make([]float64, 0),
		},
		HourlyUnits: models.Units{
			Temperature2m: "°C",
		},
	}
}

func (w *WeatherResponse) GetWeather(lat float64, lon float64) (*WeatherResponse, error) {
	fmt.Printf("GetWeather: fetching for lat=%f, lon=%f\n", lat, lon)
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=temperature_2m",
		lat, lon,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("faild to get response from weather api")
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status for response is not ok %s: ", err)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to Read the response")
	}

	if err := json.Unmarshal(body, w); err != nil {
		return nil, err
	}
	if len(w.Hourly.Temperature2m) == 0 {
		return nil, fmt.Errorf("no temperature data available for lat=%f, lon=%f", lat, lon)
	}
	fmt.Printf("GetWeather : result : %v\n" , w.Latitude )
	return w, nil

}
