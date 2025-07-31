package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/adapter"
)

func main() {
	externalService := &adapter.ExternalWeatherService{}
	weatherAdapter := &adapter.WeatherAdapter{
		ExternalWeatherService: externalService,
	}

	temp, err := weatherAdapter.GetTemperature("Tehran")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The temperature in Tehran is %.1f°C\n", temp)

}
