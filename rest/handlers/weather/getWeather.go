package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type WeatherCurrent struct {
	Temp          float64 `json:"temp"`
	FeelsLike     float64 `json:"feels_like"`
	Condition     string  `json:"condition"`
	Icon          string  `json:"icon"`
	Humidity      int     `json:"humidity"`
	WindSpeed     float64 `json:"wind_speed"`
	Precipitation float64 `json:"precipitation"`
	UVIndex       float64 `json:"uv_index"`
	Sunrise       string  `json:"sunrise"`
	Sunset        string  `json:"sunset"`
}

type ForecastItem struct {
	Time      string  `json:"time"`
	Temp      float64 `json:"temp"`
	Condition string  `json:"condition"`
	Icon      string  `json:"icon"`
}

type WeatherResponse struct {
	Location string         `json:"location"`
	Current  WeatherCurrent `json:"current"`
	Forecast []ForecastItem `json:"forecast"`
}

func (h *Handler) GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	// Trim spaces/newlines from query params
	lat := strings.TrimSpace(r.URL.Query().Get("lat"))
	lon := strings.TrimSpace(r.URL.Query().Get("lon"))

	if lat == "" || lon == "" {
		http.Error(w, "Missing coordinates", http.StatusBadRequest)
		return
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&units=metric&appid=%s",
		lat, lon, apiKey,
	)

	fmt.Println("Fetching URL:", url)

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch weather: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		http.Error(w, fmt.Sprintf("OpenWeather API error: %s - %s", resp.Status, string(body)), http.StatusInternalServerError)
		return
	}

	// Decode JSON safely into a map
	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse weather JSON: %v", err), http.StatusInternalServerError)
		return
	}

	// Extract fields safely
	location, _ := data["name"].(string)

	mainData, ok := data["main"].(map[string]interface{})
	if !ok {
		http.Error(w, "Main weather data missing", http.StatusInternalServerError)
		return
	}
	temp, _ := mainData["temp"].(float64)
	feelsLike, _ := mainData["feels_like"].(float64)
	humidityF, _ := mainData["humidity"].(float64)
	humidity := int(humidityF)

	weatherArr, _ := data["weather"].([]interface{})
	condition, icon := "", ""
	if len(weatherArr) > 0 {
		first := weatherArr[0].(map[string]interface{})
		condition, _ = first["main"].(string)
		icon, _ = first["icon"].(string)
	}

	windData, _ := data["wind"].(map[string]interface{})
	windSpeed, _ := windData["speed"].(float64)

	sysData, _ := data["sys"].(map[string]interface{})
	sunriseF, _ := sysData["sunrise"].(float64)
	sunsetF, _ := sysData["sunset"].(float64)

	weatherResp := WeatherResponse{
		Location: location,
		Current: WeatherCurrent{
			Temp:      temp,
			FeelsLike: feelsLike,
			Condition: condition,
			Icon:      icon,
			Humidity:  humidity,
			WindSpeed: windSpeed,
			Sunrise:   fmt.Sprintf("%.0f", sunriseF),
			Sunset:    fmt.Sprintf("%.0f", sunsetF),
		},
		Forecast: []ForecastItem{}, // add forecast later if needed
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherResp)
}
