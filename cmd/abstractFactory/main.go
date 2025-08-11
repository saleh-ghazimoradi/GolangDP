package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/abstractFactory"
)

func main() {
	// Create Windows UI
	windowsFactory := abstractFactory.WindowsFactory{}
	fmt.Println("Windows UI:")
	abstractFactory.CreateUI(&windowsFactory)
	fmt.Println()
	// Create macOS UI
	macOsFactory := abstractFactory.MacOsFactory{}
	fmt.Println("MacOS:")
	abstractFactory.CreateUI(&macOsFactory)
}
