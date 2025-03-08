package context

import (
	"context"
	"fmt"
	"time"
)

// Creates a context that automatically cancels after a specified duration.

func autoCancelledGoroutine(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second): // Simulating periodic work
			fmt.Println("Processing...")

		case <-ctx.Done(): // Context timeout reached
			fmt.Println("Timeout reached:", ctx.Err())
			return // Exit Goroutine
		}
	}
}

func TimeOutHelper() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	go autoCancelledGoroutine(ctx)

	time.Sleep(4 * time.Second)

}

// ✅ Best for:
// Setting timeouts for database queries and HTTP requests.
// Preventing long-running operations.
