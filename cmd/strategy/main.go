package main

import (
	"fmt"
	"github.com/saleh-ghazimoradi/GolangDP/strategy"
)

func main() {
	cart, err := strategy.NewShoppingCart(150.75, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := cart.Checkout(); err != nil {
		fmt.Println("Error:", err)
	}

	creditCard, err := strategy.NewCreditCard("1234567890123456", "123")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cart.SetPaymentStrategy(creditCard)
	if err := cart.Checkout(); err != nil {
		fmt.Println("Error:", err)
	}

	paypal, err := strategy.NewPaypal("user@example.com", "secure123")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cart.SetPaymentStrategy(paypal)
	if err := cart.Checkout(); err != nil {
		fmt.Println("Error:", err)
	}

	bitcoin, err := strategy.NewBitcoin("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	cart.SetPaymentStrategy(bitcoin)
	if err := cart.Checkout(); err != nil {
		fmt.Println("Error:", err)
	}

	_, err = strategy.NewCreditCard("1234", "12")
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = strategy.NewPaypal("invalid-email", "short")
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = strategy.NewShoppingCart(-10.0, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
