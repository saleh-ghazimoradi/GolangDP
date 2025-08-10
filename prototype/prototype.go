package prototype

// Cloneable is a prototype interface
type Cloneable interface {
	Clone() Cloneable
}

// Prototype is a concrete prototype
type Prototype struct {
	Name string
	Age  int
}

func (p *Prototype) Clone() Cloneable {
	// Deep copy to avoid shared references
	return &Prototype{
		Name: p.Name,
		Age:  p.Age,
	}
}
