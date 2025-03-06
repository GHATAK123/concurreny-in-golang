package ood

import "fmt"

// Abstraction is a fundamental concept in object-oriented programming that hides implementation details and exposes only the necessary functionality.
// In Golang, abstraction is achieved using interfaces.

// 🔹 Why Use Abstraction in Go?
// ✅ Hides implementation details and provides a clean interface.
// ✅ Makes code modular and easy to extend.
// ✅ Allows different implementations to be used interchangeably.
// ✅ Helps achieve loose coupling and flexibility.

// Logger interface (Abstraction)
// The Logger interface abstracts logging behavior without specifying implementation details.
type Logger interface {
	Log(message string)
}

// ConsoleLogger struct
type ConsoleLogger struct{}

// ConsoleLogger implements Logger interface
func (c ConsoleLogger) Log(message string) {
	fmt.Println("[Console]:", message)
}

// FileLogger struct
type FileLogger struct{}

// FileLogger implements Logger interface
func (f FileLogger) Log(message string) {
	fmt.Println("[File]: Writing to file -", message)
}

// Function that logs messages without knowing the actual implementation
// The function PerformLogging doesn't care if it's logging to a file or console.
// It relies on the Logger interface for abstraction.
func PerformLogging(logger Logger, message string) {
	logger.Log(message)
}

func Helper() {
	fmt.Println("From OOD Package!!!")
	consoleLogger := ConsoleLogger{}
	fileLogger := FileLogger{}

	// Using different loggers interchangeably
	PerformLogging(consoleLogger, "Application started")
	PerformLogging(fileLogger, "Error: Unable to open the file")
}
