package main

import "github.com/saleh-ghazimoradi/GolangDP/observer"

func main() {
	shirtItem := observer.NewItem("Nike Shirt")
	firstObserver := &observer.Customer{Id: "abc@example.com"}
	secondObserver := &observer.Customer{Id: "xyz@example.com"}

	shirtItem.Register(firstObserver)
	shirtItem.Register(secondObserver)
	shirtItem.UpdateAvailability()
}
