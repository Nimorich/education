package main

import "fmt"

// Printer interface defines methods for printing messages
type Printer interface {
	Print() string
}

// HelloWorld implements the Printer interface
type PrintHello struct {
	message string
}

// Print returns the hello world message
func (h *PrintHello) Print() string {
	return h.message
}

// NewHelloWorld creates a new HelloWorld instance
func NewPrintHello() *PrintHello {
	return &PrintHello{
		message: "Hello, World!",
	}
}

func main() {
	// Create a new HelloWorld instance
	var printer Printer = NewPrintHello()

	// Print the message
	fmt.Println(printer.Print())
}
