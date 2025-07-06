package main

import (
	"flag"
	"fmt"
	"log"
	"weather-cli/internal/weather"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(" Could not load .env file (make sure it exists in root)")
	}
	zip := flag.String("zip", "", "Zip code (e.g 500091)")
	country := flag.String("country", "IN", "Country code (default IN)")
	flag.Parse()
	if *zip == "" {
		log.Fatal("Please provide a zip code using --zip")
	}
	// apiKey := config.GetEnv("OPENWEATHERMAP_API_KEY")
	// fmt.Println("API Key loaded:", apiKey)

	lat, lon, err := weather.FetchLatLon(*zip, *country)
	if err != nil {
		log.Fatalf(" Failed to fetch coordinates: %v", err)
	}
	w, err := weather.FetchWeather(lat, lon)
	if err != nil {
		log.Fatalf("Failed to fetch weather: %v", err)
	}

	fmt.Printf("City :%s\n", w.Name)
	fmt.Printf("🌤️  Weather: %s\n", w.Description)
	fmt.Printf("🌡️  Min Temp: %.2f°C\n", w.MinTemp)
	fmt.Printf("🌡️  Max Temp: %.2f°C\n", w.MaxTemp)
	fmt.Printf("💨 Wind Speed: %.2f m/s\n", w.WindSpeed)
}
