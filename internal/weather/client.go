package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"weather-cli/internal/config"
)

// struct to read the response from the API
type WeatherResponse struct {
	Name    string `json:"name"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		TempMin float64 `json:"temp_min"`
		TempMax float64 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

// struct to send the response
type WeatherData struct {
	Name        string
	Description string
	MinTemp     float64
	MaxTemp     float64
	WindSpeed   float64
}

func FetchWeather(lat, lon float64) (WeatherData, error) {
	apiKey := config.GetEnv("OPENWEATHERMAP_API_KEY")

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%.4f&lon=%.4f&appid=%s&units=metric",
		lat, lon, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		return WeatherData{}, fmt.Errorf("error fetching weather: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return WeatherData{}, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var data WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return WeatherData{}, fmt.Errorf("failed to decode response: %v", err)
	}

	return WeatherData{
		Name:        data.Name,
		Description: data.Weather[0].Description,
		MinTemp:     data.Main.TempMin,
		MaxTemp:     data.Main.TempMax,
		WindSpeed:   data.Wind.Speed,
	}, nil
}
