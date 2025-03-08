package channels

import (
	"fmt"
	"math"
	"runtime"
	"sync"
)

func Generator(limit int) <-chan int {
	stream := make(chan int)
	go func() {
		for i := 2; i < limit; i++ {
			stream <- i
		}
		close(stream)
	}()
	return stream
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func worker(id int, numbers <-chan int, primes chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range numbers {
		if isPrime(num) {
			primes <- num
		}
	}
}

func fanOut(input <-chan int, workerCount int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(i, input, out, &wg)
	}

	// Close output channel when all workers finish
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func fanIn(primeChan <-chan int) []int {
	var primes []int
	for prime := range primeChan {
		primes = append(primes, prime)
	}
	return primes
}

func FanInFanOutConcurrenyPattern() {

	limit := 50
	workerCount := runtime.NumCPU()

	generatedChannel := Generator(limit)

	primeChannel := fanOut(generatedChannel, workerCount)

	primeNumbers := fanIn(primeChannel)

	fmt.Println(primeNumbers)

}
