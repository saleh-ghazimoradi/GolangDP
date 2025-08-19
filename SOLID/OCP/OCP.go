package OCP

// OCP states that software entities should be open for extension but closed for modification

// We can achieve this goal via interfaces and embeddings in Golang

// You can easily notice that we have extended the entity without making any modifications

type PaymentProcessor interface {
	Process(amount float64) error
}

type CreditCardPaymentProcessor struct{}

func (c *CreditCardPaymentProcessor) Process(amount float64) error {
	return nil
}

type PaypalPaymentProcessor struct{}

func (p *PaypalPaymentProcessor) Process(amount float64) error {
	return nil
}

type BitcoinPaymentProcessor struct{}

func (b *BitcoinPaymentProcessor) Process(amount float64) error {
	return nil
}

func PaymentProcess(processor PaymentProcessor, amount float64) error {
	return processor.Process(amount)
}
