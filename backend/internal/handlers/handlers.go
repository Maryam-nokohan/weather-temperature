package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"github.com/maryam-nokohan/WeatherApp/backend/internal/weather"
)

var templates = template.Must(template.ParseFiles("frontend/html/index.html", "frontend/html/city.html"))

type CityWeatherData struct {
	City        string
	Temperature string
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func SearchCityPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "unable to parse form", http.StatusBadRequest)
	}
	cityName := r.FormValue("cityName")

	if cityName == "" {
		http.Error(w, "City name is empty!", http.StatusBadRequest)
		return
	}

	city, err := weather.GetCity(cityName)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting city : %s", err), http.StatusInternalServerError)
		return
	}

	 wr := weather.NewWeatherResponse()
    wr, err = wr.GetWeather(city.Latitude, city.Longitude)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error getting weather: %v", err), http.StatusInternalServerError)
        return
    }
	
	http.Redirect(w, r, fmt.Sprintf("/city?city=%s", cityName), http.StatusSeeOther)

}

func HandleCity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	cityName := r.URL.Query().Get("city")
	if cityName == "" {
		http.Error(w, "City name is empty!", http.StatusBadRequest)
		return
	}

	city, err := weather.GetCity(cityName)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting city : %s", err), http.StatusInternalServerError)
		return
	}

	wr := weather.NewWeatherResponse()
	wr, err = wr.GetWeather(city.Latitude, city.Longitude)

	if err != nil {
		http.Error(w, fmt.Sprintf("error getting city weather : %s", err), http.StatusInternalServerError)
		return
	}

	 if len(wr.Hourly.Temperature2m) == 0 {
        http.Error(w, "No temperature data available", http.StatusInternalServerError)
        return
    }

    temperature := fmt.Sprintf("%.1f %s", wr.Hourly.Temperature2m[0], wr.HourlyUnits.Temperature2m)
    data := CityWeatherData{
        City:        cityName,
        Temperature: temperature,
    }

    if err := templates.ExecuteTemplate(w, "city.html", data); err != nil {
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    }

}
