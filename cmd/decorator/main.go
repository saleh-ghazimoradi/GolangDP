package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/decorator"
)

func main() {
	coffee := &decorator.SimpleCoffee{}
	fmt.Printf("Cost: %d, Description: %s\n", coffee.GetCost(), coffee.GetDescription())
	coffeeWithMilk := &decorator.MilkDecorator{Coffee: coffee}
	fmt.Printf("Cost: %d, Description: %s\n", coffeeWithMilk.GetCost(), coffeeWithMilk.GetDescription())
	coffeeWithMilkAndSugar := &decorator.SugarDecorator{Coffee: coffeeWithMilk}
	fmt.Printf("Cost: %d, Description: %s\n", coffeeWithMilkAndSugar.GetCost(), coffeeWithMilkAndSugar.GetDescription())

	ordinaryCoffee := decorator.OrdinaryCoffee
	cost, desc := ordinaryCoffee()
	fmt.Printf("Order: %s, Cost: $%d\n", desc, cost)

	milkyCoffee := decorator.MilkyCoffee(ordinaryCoffee)
	cost, desc = milkyCoffee()
	fmt.Printf("Order: %s, Cost: $%d\n", desc, cost)

	sweetMilkyCoffee := decorator.SweetCoffee(decorator.MilkyCoffee(decorator.SweetCoffee(decorator.OrdinaryCoffee)))
	cost, desc = sweetMilkyCoffee()
	fmt.Printf("Order: %s, Cost: $%d\n", desc, cost)

	milkySweetCoffee := decorator.MilkyCoffee(decorator.SweetCoffee(decorator.OrdinaryCoffee))
	cost, desc = milkySweetCoffee()
	fmt.Printf("Order: %s, Cost: $%d\n", desc, cost)

	// structural example
	core := &decorator.CoreHandler{}
	handler := &decorator.AuthDecorator{Handler: &decorator.LoggerDecorator{Handler: core}}
	fmt.Println(handler)

	// functional example
	query := decorator.CoreQuery
	result, err := query()
	fmt.Printf("Result: %s, Error: %v\n", result, err)

	timedQuery := decorator.TimingDecorator(decorator.CoreQuery)
	result, err = timedQuery()
	fmt.Printf("Result: %s, Error: %v\n", result, err)

	retryTimedQuery := decorator.RetryDecorator(decorator.TimingDecorator(decorator.CoreQuery), 3)
	result, err = retryTimedQuery()
	fmt.Printf("Result: %s, Error: %v\n", result, err)
}
