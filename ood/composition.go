package ood

import "fmt"

// Composition is a design principle where small, reusable components are combined to create complex structures.
// Unlike inheritance (which Golang does not support), composition allows code reusability without tightly coupling components.

// ✅ Why Use Composition Instead of Inheritance?
// Go does not support classical inheritance (no extends keyword).
// Composition promotes flexibility – behaviors can be easily changed.
// Loosely coupled code – modifying one struct doesn’t impact others.
// Better reusability – functionality can be shared across multiple types.

type Engine interface {
	Start()
}

// PetrolEngine struct
type PetrolEngine struct{}

func (p PetrolEngine) Start() {
	fmt.Println("Petrol Engine started")
}

// ElectricEngine struct
type ElectricEngine struct{}

func (e ElectricEngine) Start() {
	fmt.Println("Electric Engine started silently")
}

// Car struct with Engine interface
type Car struct {
	Engine Engine
	Brand  string
}

func CompositionHelper() {
	petrolCar := Car{Engine: PetrolEngine{}, Brand: "Ford"}
	electricCar := Car{Engine: ElectricEngine{}, Brand: "Tesla"}

	petrolCar.Engine.Start()   // Petrol Engine started
	electricCar.Engine.Start() // Electric Engine started silently
	// Go supports interface-based composition, where different implementations can be swapped.
}

// ✅ Composition is preferred over inheritance in Go.
// ✅ Struct embedding allows code reuse without inheritance.
// ✅ Interfaces + Composition = High flexibility.
// ✅ Multi-level composition allows more modularity.
