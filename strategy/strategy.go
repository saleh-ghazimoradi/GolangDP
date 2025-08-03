package strategy

import (
	"fmt"
	"strings"
)

type PaymentStrategy interface {
	Pay(amount float64) error
}

type CreditCard struct {
	CardNumber string
	CVV2       string
}

func (c *CreditCard) Pay(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: must be positive")
	}
	fmt.Printf("Paying %.2f using Credit Card %s\n", amount, c.CardNumber)
	return nil
}

func NewCreditCard(cardNumber, cvv2 string) (*CreditCard, error) {
	if len(cardNumber) < 16 || len(cardNumber) > 19 {
		return nil, fmt.Errorf("invalid card number: must be 16-19 digits")
	}
	if len(cvv2) != 3 && len(cvv2) != 4 {
		return nil, fmt.Errorf("invalid CVV2: must be 3 or 4 digits")
	}
	return &CreditCard{
		CardNumber: cardNumber,
		CVV2:       cvv2,
	}, nil
}

type Paypal struct {
	Email    string
	Password string
}

func (p *Paypal) Pay(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: must be positive")
	}
	fmt.Printf("Paying %.2f using PayPal %s\n", amount, p.Email)
	return nil
}

func NewPaypal(email, password string) (*Paypal, error) {
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return nil, fmt.Errorf("invalid email address")
	}
	if len(password) < 6 {
		return nil, fmt.Errorf("invalid password: must be at least 6 characters")
	}
	return &Paypal{
		Email:    email,
		Password: password,
	}, nil
}

type Bitcoin struct {
	Address string
}

func (b *Bitcoin) Pay(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: must be positive")
	}
	fmt.Printf("Paying %.2f using Bitcoin %s\n", amount, b.Address)
	return nil
}

func NewBitcoin(address string) (*Bitcoin, error) {
	if len(address) < 26 || len(address) > 35 {
		return nil, fmt.Errorf("invalid Bitcoin address: must be 26-35 characters")
	}
	return &Bitcoin{
		Address: address,
	}, nil
}

type ShoppingCart struct {
	paymentStrategy PaymentStrategy
	amount          float64
}

func (s *ShoppingCart) SetPaymentStrategy(paymentStrategy PaymentStrategy) {
	s.paymentStrategy = paymentStrategy
}

func (s *ShoppingCart) Checkout() error {
	if s.paymentStrategy == nil {
		return fmt.Errorf("payment strategy is nil")
	}
	return s.paymentStrategy.Pay(s.amount)
}

func NewShoppingCart(amount float64, paymentStrategy PaymentStrategy) (*ShoppingCart, error) {
	if amount <= 0 {
		return nil, fmt.Errorf("invalid amount: must be positive")
	}
	return &ShoppingCart{
		paymentStrategy: paymentStrategy,
		amount:          amount,
	}, nil
}
