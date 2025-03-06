package ood

import "fmt"

// Encapsulation is one of the key principles of Object-Oriented Programming (OOP), which helps in data hiding and controlling access to the internal state of an object.

// 🔥 Encapsulation in Golang:
// Go achieves encapsulation using struct and methods.
// Access control is done using visibility rules:
// Exported (Public): Names starting with uppercase letters (User, GetBalance)
// Unexported (Private): Names starting with lowercase letters (password, balance)

// Encapsulated struct
type Employee struct {
	name   string
	salary float64
}

// Constructor function
func NewEmployee(name string, salary float64) *Employee {
	return &Employee{name: name, salary: salary}
}

// Getter method
func (e *Employee) GetSalary() float64 {
	return e.salary
}

// Setter method with validation
func (e *Employee) SetSalary(newSalary float64) {
	if newSalary > 0 {
		e.salary = newSalary
	} else {
		fmt.Println("Invalid salary amount")
	}
}

func EncapsulationHelper() {

	emp := NewEmployee("John Doe", 50000)

	fmt.Println("Initial Salary:", emp.GetSalary())

	emp.SetSalary(60000) // Valid update
	fmt.Println("Updated Salary:", emp.GetSalary())

	emp.SetSalary(-10000) // Invalid update

}
