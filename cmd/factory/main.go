package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/factory"
)

func main() {
	pizza, _ := factory.GetPizzaProduct("Iranian")
	fmt.Printf("%+v\n", pizza)
}
