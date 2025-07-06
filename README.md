# 🌤️ Weather CLI (Go)

A minimal command-line application written in Go that fetches and displays current weather data using the [OpenWeatherMap API](https://openweathermap.org/api).

---

## 📦 Features

- Fetch weather using `--zip` and `--country`
- Shows:
  - City name
  - Weather description
  - Min & Max temperature
  - Wind speed
- `--pretty` flag for formatted output
- Loads API key from `.env` file
- Clean Go module structure

---

## 🚀 Installation

```bash
git clone https://github.com/YOUR_USERNAME/weather-cli.git
cd weather-cli
go mod tidy
