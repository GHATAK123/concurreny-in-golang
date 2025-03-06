package ood

import "fmt"

type PaymentMethod interface {
	Pay(amount float64)
}

// CreditCard struct
type CreditCard struct {
	CardNumber string
}

// Pay method for CreditCard
func (c CreditCard) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card: %s\n", amount, c.CardNumber)
}

// PayPal struct
type PayPal struct {
	Email string
}

// Pay method for PayPal
func (p PayPal) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal: %s\n", amount, p.Email)
}

// Function demonstrating polymorphism
func ProcessPayment(p PaymentMethod, amount float64) {
	p.Pay(amount)
}

func PolyHelper() {
	cc := CreditCard{CardNumber: "1234-5678-9012-3456"}
	pp := PayPal{Email: "user@example.com"}

	ProcessPayment(cc, 100.50)
	ProcessPayment(pp, 75.00)
}
