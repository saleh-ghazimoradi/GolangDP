package ISP

// ISP states that clients should never be forced to depend on the interfaces which they don't use

// This principle encourages us to use smaller interfaces yet more focused

// This example below violates the ISP

type Document struct{}

type Printers interface {
	Print(d Document) error
	Scan(d Document) error
	Fax(d Document) error
}

type SimplePrinter struct{}

func (p *SimplePrinter) Print(d Document) error {
	return nil
}

func (p *SimplePrinter) Scan(d Document) error {
	return nil
}

func (p *SimplePrinter) Fax(d Document) error {
	return nil
}

// Instead of creating a big interface forcing the clients to implement, creat smaller and more focused like the ones below

type Printer interface {
	Print(d Document) error
}

type Scanner interface {
	Scan(d Document) error
}

type Faxer interface {
	Fax(d Document) error
}

type OrdinaryPrinter struct{}

func (o *OrdinaryPrinter) Print(d Document) error {
	return nil
}

type MultifunctionPrinter struct{}

func (m *MultifunctionPrinter) Print(d Document) error {
	return nil
}

func (m *MultifunctionPrinter) Scan(d Document) error {
	return nil
}

func (m *MultifunctionPrinter) Fax(d Document) error {
	return nil
}
