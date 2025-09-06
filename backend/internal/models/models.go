package models

type Hourly struct {
    Time          []string  `json:"time"`
    Temperature2m []float64 `json:"temperature_2m"`
}

type Units struct {
    Temperature2m string `json:"temperature_2m"`
}

type City struct {
    CountryName string  `json:"country"`
    CityName     string  `json:"name"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}
    
type GeoResponse struct {
    Results []City `json:"results"`
}
