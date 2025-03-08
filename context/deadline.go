package context

import (
	"context"
	"fmt"
	"time"
)

func autoCancelledGoroutineDeadline(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second): // Simulating periodic work
			fmt.Println("Processing...")

		case <-ctx.Done(): // Context deadline reached
			fmt.Println("Deadline reached:", ctx.Err())
			return // Exit Goroutine
		}
	}
}

func DeadlineHelper() {
	// Set an exact deadline (3 seconds from now)
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // Ensure cleanup

	go autoCancelledGoroutineDeadline(ctx)

	time.Sleep(5 * time.Second)

}
