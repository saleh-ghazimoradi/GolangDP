package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/builder"
)

func main() {
	buick := builder.NewCarBuilder().
		Make("Buick").
		Model("Skylark").
		Color("white").
		Seats(5).
		Engine("V8").
		Transmission("Automatic").
		Build()
	fmt.Printf("%+v\n", buick)
}
