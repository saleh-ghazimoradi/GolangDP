package observer

import "fmt"

// Subject is the subject of observer pattern
type Subject interface {
	Register(observer Observer)
	Deregister(observer Observer)
	NotifyAll()
}

// Item is the concrete subject of observer pattern
type Item struct {
	ObserverList []Observer
	Name         string
	InStock      bool
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is in stock\n", i.Name)
	i.InStock = true
	i.NotifyAll()
}

func (i *Item) Register(observer Observer) {
	i.ObserverList = append(i.ObserverList, observer)
}

func (i *Item) Deregister(observer Observer) {
	for j, obs := range i.ObserverList {
		if obs == observer {
			i.ObserverList = append(i.ObserverList[:j], i.ObserverList[j+1:]...)
			break
		}
	}
}

func (i *Item) NotifyAll() {
	for _, observer := range i.ObserverList {
		observer.Update(i.Name)
	}
}

func NewItem(name string) *Item {
	return &Item{
		Name: name,
	}
}

// Observer is the observer of observer pattern
type Observer interface {
	Update(string)
	GetId() string
}

// Customer is the concrete observer of observer pattern
type Customer struct {
	Id string
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.Id, itemName)
}

func (c *Customer) GetId() string {
	return c.Id
}
