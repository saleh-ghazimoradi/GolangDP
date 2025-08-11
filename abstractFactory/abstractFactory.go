package abstractFactory

import "fmt"

// Abstract Products
type Button interface {
	Render() string
}

type Checkbox interface {
	Check() string
}

// Concrete Products for Windows
type WindowsButton struct{}

func (w *WindowsButton) Render() string {
	return "Rendering windows button"
}

type WindowsCheckbox struct{}

func (w *WindowsCheckbox) Check() string {
	return "Checking windows checkbox"
}

// Concrete Products for macOS
type MacOsButton struct{}

func (m *MacOsButton) Render() string {
	return "Rendering macOS button"
}

type MacOsCheckbox struct{}

func (m *MacOsCheckbox) Check() string {
	return "Checking macOS checkbox"
}

// Abstract Factory
type UIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Concrete Factory for Windows
type WindowsFactory struct{}

func (w *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (w *WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

// Concrete Factory for macOS
type MacOsFactory struct{}

func (m *MacOsFactory) CreateButton() Button {
	return &MacOsButton{}
}

func (m *MacOsFactory) CreateCheckbox() Checkbox {
	return &MacOsCheckbox{}
}

// Client code
func CreateUI(factory UIFactory) {
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()
	fmt.Println(button.Render())
	fmt.Println(checkbox.Check())
}
