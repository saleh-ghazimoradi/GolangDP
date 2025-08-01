package decorator

// struct based decorator
type Coffee interface {
	GetCost() int
	GetDescription() string
}

type SimpleCoffee struct{}

func (s *SimpleCoffee) GetCost() int {
	return 5
}

func (s *SimpleCoffee) GetDescription() string {
	return "Simple Coffee"
}

type MilkDecorator struct {
	Coffee Coffee
}

func (m *MilkDecorator) GetCost() int {
	return m.Coffee.GetCost() + 2
}

func (m *MilkDecorator) GetDescription() string {
	return m.Coffee.GetDescription() + ",milk"
}

type SugarDecorator struct {
	Coffee Coffee
}

func (s *SugarDecorator) GetCost() int {
	return s.Coffee.GetCost() + 1
}

func (s *SugarDecorator) GetDescription() string {
	return s.Coffee.GetDescription() + ",sugar"
}

// functional based decorator peculiar to Golang
type CoffeeFunc func() (cost int, description string)

func OrdinaryCoffee() (int, string) {
	return 5, "Ordinary Coffee"
}

func MilkyCoffee(coffee CoffeeFunc) CoffeeFunc {
	return func() (int, string) {
		cost, description := coffee()
		return cost + 2, description + ", milk"
	}
}

func SweetCoffee(coffee CoffeeFunc) CoffeeFunc {
	return func() (int, string) {
		cost, description := coffee()
		return cost + 1, description + ", sugar"
	}
}
