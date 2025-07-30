package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/factory"
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
}
