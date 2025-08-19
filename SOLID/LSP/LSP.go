package LSP

import (
	"fmt"
	"math"
)

// LSP states that the objects of a superclass must be replaceable with the objects of a subclass

// This principle is implemented via interfaces in Golang

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Print(s Shape) {
	fmt.Println(s.Area())
}
