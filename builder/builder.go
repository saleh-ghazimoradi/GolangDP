package builder

type Car struct {
	Make         string
	Model        string
	Color        string
	Seats        int
	Engine       string
	Transmission string
}

type CarBuilder struct {
	car *Car
}

func (b *CarBuilder) Make(make string) *CarBuilder {
	b.car.Make = make
	return b
}

func (b *CarBuilder) Model(model string) *CarBuilder {
	b.car.Model = model
	return b
}

func (b *CarBuilder) Color(color string) *CarBuilder {
	b.car.Color = color
	return b
}

func (b *CarBuilder) Seats(seats int) *CarBuilder {
	b.car.Seats = seats
	return b
}

func (b *CarBuilder) Engine(engine string) *CarBuilder {
	b.car.Engine = engine
	return b
}

func (b *CarBuilder) Transmission(transmission string) *CarBuilder {
	b.car.Transmission = transmission
	return b
}

func (b *CarBuilder) Build() *Car {
	return b.car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{
		car: &Car{},
	}
}
