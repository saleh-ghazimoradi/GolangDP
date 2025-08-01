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
}
