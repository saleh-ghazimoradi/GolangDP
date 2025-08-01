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
}
