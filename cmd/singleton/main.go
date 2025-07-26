package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/singleton"
)

func main() {
	s := singleton.GetInstance()
	s.SetData("Hello World")
	fmt.Println(s.GetData())
}
