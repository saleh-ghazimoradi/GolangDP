package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/observer"
)

func main() {
	shirtItem := observer.NewItem("Nike Shirt")
	firstObserver := &observer.Customer{Id: "abc@example.com"}
	secondObserver := &observer.Customer{Id: "xyz@example.com"}

	shirtItem.Register(firstObserver)
	shirtItem.Register(secondObserver)
	shirtItem.UpdateAvailability()
	fmt.Println("---------------------------------")

	// News Feed System
	agency := &observer.NewsAgency{}
	app1 := &observer.MobileApp{Name: "App1"}
	app2 := &observer.MobileApp{Name: "App2"}
	website := &observer.Website{Url: "example.com"}
	agency.Register(app1)
	agency.Register(app2)
	agency.Register(website)
	agency.Publish("Breaking News: Go 2.0 Released!")
	agency.Deregister(app2)
	fmt.Println("\nPublishing another article...")
	agency.Publish("Tech: New AI Breakthrough")
}
