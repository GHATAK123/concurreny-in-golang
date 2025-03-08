package context

import (
	"context"
	"fmt"
	"time"
)

// The context.WithValue(ctx, key, value) function allows you to pass request-scoped values through the context
// The context.WithValue() function is useful for passing metadata (like user IDs or request tracking information) across Goroutines.

type contextKey string

func worker(ctx context.Context) {
	// Simulate some processing time
	time.Sleep(1 * time.Second)

	// Retrieve the user ID from the context
	key := contextKey("userID")
	if userID, ok := ctx.Value(key).(int); ok {
		fmt.Println("Worker received User ID from context:", userID)
	} else {
		fmt.Println("Worker: No User ID found in context")
	}
}

func ValueHelper() {
	// Create a base context
	ctx := context.Background()

	// Define a key and store user metadata in context
	key := contextKey("userID")
	ctx = context.WithValue(ctx, key, 12345) // Store user ID

	// Start a worker Goroutine that uses the context
	go worker(ctx)

	// Give the Goroutine some time to finish before exiting main()
	time.Sleep(2 * time.Second)

}
