package factory

import (
	"errors"
	"fmt"
)

type PaymentProcess interface {
	Payment(amount float64) error
}

type paymentProcess struct {
	apiKey   string
	currency string
}

func (p *paymentProcess) Payment(amount float64) error {
	if p.apiKey == "" {
		return errors.New("PayPal payment failed: missing API key")
	}
	fmt.Printf("Processing PayPal payment of %.2f %s\n", amount, p.currency)
	return nil
}

type paypal struct {
	paymentProcess
}

func NewPaypal(apiKey, currency string) PaymentProcess {
	return &paypal{
		paymentProcess{
			apiKey:   apiKey,
			currency: currency,
		},
	}
}

type stripe struct {
	paymentProcess
}

func NewStripe(apiKey, currency string) PaymentProcess {
	return &stripe{
		paymentProcess{
			apiKey:   apiKey,
			currency: currency,
		},
	}
}

type Crypto struct {
	paymentProcess
}

func NewCrypto(apiKey, currency string) PaymentProcess {
	return &Crypto{
		paymentProcess{
			apiKey:   apiKey,
			currency: currency,
		},
	}
}

type Config struct {
	Type     string
	APIKey   string
	Currency string
}

func NewPaymentProcess(cfg Config) (PaymentProcess, error) {
	switch cfg.Type {
	case "Paypal":
		return NewPaypal(cfg.APIKey, cfg.Currency), nil
	case "Stripe":
		return NewStripe(cfg.APIKey, cfg.Currency), nil
	case "Crypto":
		return NewCrypto(cfg.APIKey, cfg.Currency), nil
	default:
		return nil, fmt.Errorf("unknown payment process type: %s", cfg.Type)
	}
}
