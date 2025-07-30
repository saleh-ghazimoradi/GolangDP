package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/factory"
	"log"
)

func main() {
	cfg := factory.PizzaConfig{Type: "Iranian", Name: "Garlic & Steak", Price: 15}
	pizza, err := factory.NewPizza(cfg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Pizza: %s, Price: $%.2f, Ingredients: %v\n", pizza.Name(), pizza.Price(), pizza.Ingredients())

	cfg = factory.PizzaConfig{Type: "Italian", Name: "Margherita", Price: 12.0}
	pizza, err = factory.NewPizza(cfg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Pizza: %s, Price: $%.2f, Ingredients: %v\n", pizza.Name(), pizza.Price(), pizza.Ingredients())

	configs := []factory.Config{
		{
			Type:     "Paypal",
			APIKey:   "paypal-api-key-123",
			Currency: "USD",
		},
		{
			Type:     "Stripe",
			APIKey:   "stripe-api-key-456",
			Currency: "EUR",
		},
		{
			Type:     "Crypto",
			APIKey:   "crypto-api-key-789",
			Currency: "BTC",
		},
		{
			Type:     "Paypal",
			APIKey:   "",
			Currency: "USD",
		},
		{
			Type:     "Unknown",
			APIKey:   "invalid-key",
			Currency: "USD",
		},
	}

	for _, cfg := range configs {
		fmt.Printf("\nAttempting to create payment processor for type: %s\n", cfg.Type)

		paymentProcessor, err := factory.NewPaymentProcess(cfg)
		if err != nil {
			log.Printf("Error creating payment processor: %v", err)
			continue
		}

		amount := 100.0
		err = paymentProcessor.Payment(amount)
		if err != nil {
			log.Printf("Payment failed: %v", err)
		} else {
			fmt.Printf("Payment of %.2f %s processed successfully\n", amount, cfg.Currency)
		}
	}
}
