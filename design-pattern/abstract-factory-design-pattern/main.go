package main

import "fmt"

// Abstract Product: Button
type Button interface {
	Paint()
}

// Abstract Product: Checkbox
type Checkbox interface {
	Check()
}

// Concrete Product: WindowsButton
type WindowsButton struct{}

func (wb *WindowsButton) Paint() {
	fmt.Println("Painting a Windows button.")
}

// Concrete Product: WindowsCheckbox
type WindowsCheckbox struct{}

func (wc *WindowsCheckbox) Check() {
	fmt.Println("Checking a Windows checkbox.")
}

// Concrete Product: MacOSButton
type MacOSButton struct{}

func (mb *MacOSButton) Paint() {
	fmt.Println("Painting a MacOS button.")
}

// Concrete Product: MacOSCheckbox
type MacOSCheckbox struct{}

func (mc *MacOSCheckbox) Check() {
	fmt.Println("Checking a MacOS checkbox.")
}

// Abstract Factory: GUIFactory
type GUIFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// Concrete Factory: WindowsFactory
type WindowsFactory struct{}

func (wf *WindowsFactory) CreateButton() Button {
	return &WindowsButton{}
}

func (wf *WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{}
}

// Concrete Factory: MacOSFactory
type MacOSFactory struct{}

func (mf *MacOSFactory) CreateButton() Button {
	return &MacOSButton{}
}

func (mf *MacOSFactory) CreateCheckbox() Checkbox {
	return &MacOSCheckbox{}
}

// Client code that works with factories and products
func main() {
	var factory GUIFactory

	// Choose a factory at runtime (can be dynamic based on environment)
	operatingSystem := "windows" // Change to "mac" for MacOS

	if operatingSystem == "windows" {
		factory = &WindowsFactory{}
	} else if operatingSystem == "mac" {
		factory = &MacOSFactory{}
	}

	// Use the factory to create products
	button := factory.CreateButton()
	checkbox := factory.CreateCheckbox()

	// Interact with the products through their interfaces
	button.Paint()
	checkbox.Check()
}
