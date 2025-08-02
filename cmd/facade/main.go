package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/facade"
	"log"
)

func main() {
	homeTheater := facade.NewHomeTheaterFacade()
	homeTheater.WatchMovie()
	fmt.Println()
	homeTheater.EndMovie()
	fmt.Println()
	fmt.Println("Real world example")
	walletFacade := facade.NewWalletFacade("Admin", 1111)
	fmt.Println()
	err := walletFacade.AddMoneyToWallet("Admin", 1111, 10)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("Admin", 1111, 5)
	if err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
}
