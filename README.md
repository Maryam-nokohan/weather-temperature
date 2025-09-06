# 🌤️ WeatherApp

WeatherApp is a simple educational weather lookup application built with **Go** (backend) and a **HTML/CSS frontend**.  
It allows users to search for a city, fetch real-time weather data, and display the results in a clean interface.  

---

## 📂 Project Structure

```bash
.
├── backend
│   ├── api
│   │   ├── api.go
│   │   └── health.go
│   ├── cmd
│   │   └── main.go
│   ├── config
│   │   └── config.go
│   ├── docs
│   │   ├── coordinate_sample.json
│   │   └── weather_sample.json
│   └── internal
│       ├── handlers
│       │   └── handlers.go
│       ├── models
│       │   └── models.go
│       └── weather
│           ├── cities.go
│           └── weather.go
├── frontend
│   ├── css
│   │   ├── cityStyle.css
│   │   └── style.css
│   └── html
│       ├── 404.html
│       ├── city.html
│       └── index.html
├── go.mod
└── go.sum

```


---

## 🚀 Features

- Search weather by city name  
- Fetch weather data from open meteo API  
- Simple responsive frontend (HTML + CSS)  
- Go backend using [chi router](https://github.com/go-chi/chi)  
- Config management via `.env`  
- Error handling with custom error page (`404.html`)  
- Example weather and coordinate data in `backend/docs`  

---

## 🛠️ Installation & Setup

### 1. Clone the repo

```bash
git clone https://github.com/Maryam-nokohan/weather-temperature.git

cd weather-temperature
```
### 2. Install dependencies

```bash
go mod tidy
```

### 3. Environment variables

Create a .env file in the project root:

```bash
ADDR=:8080
```

1. Run the app
```bash
go run backend/cmd/main.go
```

1. Open in browser

Go to `http://localhost:8080`

# 📖 Tech Stack

Backend: Go, chi router

Frontend: HTML, CSS

Templates: Go’s html/template

Config: godotenv

# 🤝 Contributing

Fork this repo

Create a feature branch (git checkout -b feature/my-feature)

Commit your changes (git commit -m "feat: add my feature")

Push to your fork (git push origin feature/my-feature)

Open a Pull Request

📜 License

MIT License – feel free to use this project as you like.

---
