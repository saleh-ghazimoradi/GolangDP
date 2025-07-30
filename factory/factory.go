package factory

import "fmt"

type PizzaProduct interface {
	SetName(name string)
	SetPrice(price float64)
	GetName() string
	GetPrice() float64
}

type ConcretePizzaProduct struct {
	name  string
	price float64
}

func (p *ConcretePizzaProduct) SetName(name string) {
	p.name = name
}

func (p *ConcretePizzaProduct) SetPrice(price float64) {
	p.price = price
}

func (p *ConcretePizzaProduct) GetName() string {
	return p.name
}

func (p *ConcretePizzaProduct) GetPrice() float64 {
	return p.price
}

type ItalianPizzaProduct struct {
	ConcretePizzaProduct
}

func NewItalianPizzaProduct(name string, price float64) PizzaProduct {
	return &ConcretePizzaProduct{
		name:  name,
		price: price,
	}
}

type AmericanPizzaProduct struct {
	ConcretePizzaProduct
}

func NewAmericanPizzaProduct(name string, price float64) PizzaProduct {
	return &ConcretePizzaProduct{
		name:  name,
		price: price,
	}
}

type IranianPizzaProduct struct {
	ConcretePizzaProduct
}

func NewIranianPizzaProduct(name string, price float64) PizzaProduct {
	return &ConcretePizzaProduct{
		name:  name,
		price: price,
	}
}

func GetPizzaProduct(pizzaType string) (PizzaProduct, error) {
	switch pizzaType {
	case "Iranian":
		return NewIranianPizzaProduct("Garlic & Steak", 15), nil
	case "Italian":
		return NewItalianPizzaProduct("Margherita", 12), nil
	case "American":
		return NewAmericanPizzaProduct("Cake Pizza", 10), nil
	default:
		return nil, fmt.Errorf("pizza type %s is not supported", pizzaType)
	}
}
