package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/prototype"
)

func main() {
	original := &prototype.Prototype{Name: "John", Age: 30}
	clone := original.Clone().(*prototype.Prototype)
	fmt.Println(original)
	fmt.Println(clone)
}
