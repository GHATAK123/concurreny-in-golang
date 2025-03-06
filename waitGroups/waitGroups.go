package waitgroups

import (
	"fmt"
	"sync"
)

/* A WaitGroup waits for a collection of goroutines to finish.
The main goroutine calls Add on the WaitGroup object to set the number of goroutines to wait for.
Every goroutine takes a pointer to the WaitGroup value as a parameter when it is invoked. When each of the goroutines runs, it calls Done when finished.
The main() calls the Wait() method to block itself until all the goroutines have finished. */

func heavyFunction1(wg *sync.WaitGroup) {
	defer wg.Done()
	// Do a lot of stuff
	fmt.Println("HeavyFunction1")
}

func heavyFunction2(wg *sync.WaitGroup) {
	defer wg.Done()
	// Do a lot of stuff
	fmt.Println("HeavyFunction2")
}

func Helper() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go heavyFunction1(wg)
	go heavyFunction2(wg)
	wg.Wait()
	fmt.Println("All Finished!")
}
