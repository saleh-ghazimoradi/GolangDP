package adapter

import (
	"strconv"
	"strings"
)

// WeatherProvider is the interface our application expects
type WeatherProvider interface {
	GetTemperature(city string) (float64, error)
}

// ExternalWeatherService represents a third-party weather service
type ExternalWeatherService struct{}

func (e *ExternalWeatherService) GetWeatherData(location string) (string, error) {
	// Simulating an API call to the external service
	return "16.5°C", nil
}

// WeatherAdapter adapts ExternalWeatherService to WeatherProvider
type WeatherAdapter struct {
	ExternalWeatherService *ExternalWeatherService
}

func (w *WeatherAdapter) GetTemperature(city string) (float64, error) {
	data, err := w.ExternalWeatherService.GetWeatherData(city)
	if err != nil {
		return 0, err
	}
	// Convert the string temperature to float64
	temp, err := strconv.ParseFloat(strings.TrimSuffix(data, "°C"), 64)
	if err != nil {
		return 0, err
	}

	return temp, nil
}
