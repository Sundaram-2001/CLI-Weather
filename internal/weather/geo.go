package weather

import (
	"encoding/json"
	"fmt"
	"net/http"

	// "net/http"
	// "github.com/joho/godotenv"
	// "os"
	"weather-cli/internal/config"
)

type GeocodeResponse struct {
	Zip     string  `json:"zip"`
	Name    string  `json:"name"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
}

func FetchLatLon(zip, country string) (float64, float64, error) {
	apiKey := config.GetEnv("OPENWEATHERMAP_API_KEY")
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/zip?zip=%s,%s&appid=%s", zip, country, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("API responded with status %d", resp.StatusCode)
	}
	var data GeocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, 0, fmt.Errorf("failed to decode response: %v", err)
	}

	return data.Lat, data.Lon, nil
}
