# 🌤️ Weather CLI (Go)

A minimal command-line application built in Go that fetches and displays real-time weather information using the [OpenWeatherMap API](https://openweathermap.org/api). It takes a ZIP code and country code as input and shows you the current weather, temperature, and wind speed — right in your terminal.

---

## 📦 Features

- 🔐 Loads API key securely from `.env`
- 🌍 Supports ZIP and country code as input (`--zip`, `--country`)
- 📍 Converts ZIP to lat/lon using OpenWeather Geocoding API
- 🌤️ Displays:
  - City name
  - Weather description
  - Min/Max temperature (°C)
  - Wind speed
- 🧪 Simple CLI flag parsing using Go’s `flag` package
- 🧱 Clean modular Go project structure

---

## 🚀 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/YOUR_USERNAME/weather-cli.git
cd weather-cli
