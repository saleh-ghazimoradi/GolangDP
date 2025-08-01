package decorator

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
