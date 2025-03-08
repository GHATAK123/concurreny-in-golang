package context

import (
	"context"
	"fmt"
	"time"
)

// context.WithCancel(ctx)
// Creates a cancelable context.
// Calling cancel() notifies all derived goroutines to stop execution.

func manuallyCancelledGoroutine(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Gouroutine Cancelled...")
			return
		default:
			fmt.Println("Goroutine running...")
			time.Sleep(time.Millisecond * 500)
		}
	}

}

func CancelHelper() {

	ctx, cancel := context.WithCancel(context.Background())

	go manuallyCancelledGoroutine(ctx)

	time.Sleep(2 * time.Second)
	cancel() // Cancel the context maunally after 2 sec
	time.Sleep(1 * time.Second)

}

// Gracefully stopping long-running goroutines.
// Avoiding leaked goroutines.
