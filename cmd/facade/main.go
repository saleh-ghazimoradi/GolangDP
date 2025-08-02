package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/facade"
)

func main() {
	homeTheater := facade.NewHomeTheaterFacade()
	homeTheater.WatchMovie()
	fmt.Println()
	homeTheater.EndMovie()
}
