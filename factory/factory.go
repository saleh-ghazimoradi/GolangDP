package factory

import "fmt"

type Pizza interface {
	Name() string
	Price() float64
	Ingredients() []string
}

type BasePizza struct {
	name        string
	price       float64
	ingredients []string
}

func (b *BasePizza) Name() string {
	return b.name
}

func (b *BasePizza) Price() float64 {
	return b.price
}

func (b *BasePizza) Ingredients() []string {
	return b.ingredients
}

type IranianPizza struct {
	BasePizza
}

func NewIranianPizza(name string, price float64) Pizza {
	return &IranianPizza{BasePizza{name, price, []string{"Garlic", "Steak", "Mushroom"}}}
}

type ItalianPizza struct {
	BasePizza
}

func NewItalianPizza(name string, price float64) Pizza {
	return &ItalianPizza{BasePizza{name, price, []string{"Tomato Sauce", "Mozzarella", "Basil"}}}
}

type AmericaPizza struct {
	BasePizza
}

func NewAmericaPizza(name string, price float64) Pizza {
	return &AmericaPizza{
		BasePizza{name, price, []string{"Pepperoni", "Cheddar", "BBQ Sauce"}},
	}
}

type PizzaConfig struct {
	Type  string
	Name  string
	Price float64
}

func NewPizza(cfg PizzaConfig) (Pizza, error) {
	switch cfg.Type {
	case "Iranian":
		return NewIranianPizza(cfg.Name, cfg.Price), nil
	case "America":
		return NewAmericaPizza(cfg.Name, cfg.Price), nil
	case "Italian":
		return NewItalianPizza(cfg.Name, cfg.Price), nil
	default:
		return nil, fmt.Errorf("unknown type: %s", cfg.Type)
	}
}
