package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/maryam-nokohan/WeatherApp/backend/internal/models"
)


type WeatherResponse struct {
    Latitude             float64 `json:"latitude"`
    Longitude            float64 `json:"longitude"`
    Elevation            float64 `json:"elevation"`
    Timezone             string  `json:"timezone"`
    TimezoneAbbreviation string  `json:"timezone_abbreviation"`
    Hourly               models.Hourly  `json:"hourly"`
    HourlyUnits          models.Units   `json:"hourly_units"`
}

type Weather interface {
	GetWeather(lat float64 , lon float64)(*WeatherResponse , error)
}

