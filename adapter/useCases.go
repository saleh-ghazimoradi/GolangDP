package adapter

import "fmt"

// PaymentProcessor is what the client expects
type PaymentProcessor interface {
	ProcessPayment(amount float64, customerId string) (string, error)
}

// Stripe payment system Adaptee
type Stripe struct{}

func (s *Stripe) Charge(amount float64, user string) (string, error) {
	return fmt.Sprintf("Stripe charged $%.2f for user %s", amount, user), nil
}

// StripeAdapter Adapts Stripe to PaymentProcessor
type StripeAdapter struct {
	Stripe *Stripe
}

func (s *StripeAdapter) ProcessPayment(amount float64, customerId string) (string, error) {
	return s.Stripe.Charge(amount, customerId)
}

// PayPal payment system Adaptee
type PayPal struct{}

func (p *PayPal) MakePayment(user string, amountInCents int) (string, error) {
	return fmt.Sprintf("PayPal paid %d cents for user %s", amountInCents, user), nil
}

// PayPalAdapter adapts PayPal to PaymentProcessor
type PayPalAdapter struct {
	PayPal *PayPal
}

func (s *PayPalAdapter) ProcessPayment(amount float64, customerId string) (string, error) {
	amountInCents := int(amount * 100)
	return s.PayPal.MakePayment(customerId, amountInCents)
}
