package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/adapter"
)

func process(p adapter.PaymentProcessor, amount float64, customerId string) {
	result, err := p.ProcessPayment(amount, customerId)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Payment result:", result)
}

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

	stripe := &adapter.Stripe{}
	stripeAdapter := &adapter.StripeAdapter{Stripe: stripe}
	process(stripeAdapter, 100.5, "user123")

	paypal := &adapter.PayPal{}
	paypalAdapter := &adapter.PayPalAdapter{PayPal: paypal}
	process(paypalAdapter, 100.5, "user123")

}
