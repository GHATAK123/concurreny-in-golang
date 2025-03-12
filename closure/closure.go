package closure

import (
	"fmt"
	"time"
)

// A closure is an anonymous function that captures and remembers the variables from its surrounding scope, even after the scope has exited.

func Helper() {
	fmt.Println("From Closure Package!!!")
	basicClosure()
	fmt.Println(closureEithState()())
	closureWithGoroutine()

}

func basicClosure() {
	message := "surrounding variable out of scope"

	func() {
		fmt.Println(message) // Closure captures 'message' variable
	}()

	// 	✅ Explanation:
	// The anonymous function inside basicClosure accesses the message variable from the outer function.
	// Even though message is declared outside the function, the closure remembers its value.
}

func closureEithState() func() int {
	count := 10
	return func() int {
		count++
		return count
	}
}

func closureWithGoroutine() {
	for i := 1; i <= 3; i++ {
		go func(num int) {
			fmt.Printf("%d ", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println()
	// 	✅ Explanation:
	// The closure explicitly captures i as a function argument, avoiding race conditions.
	// If i was directly used inside the closure, all goroutines would print the last value of i.
}
