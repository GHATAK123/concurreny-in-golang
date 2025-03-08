package channels

import (
	"fmt"
	"time"
)

// UnBuffered Channel | synchronus | Blocking
// the send operation does not complete until there is a receiver to accept the value.
// One can think of an unbuffered channel as if there is no space in the channel for data.
// There must be a receiver ready to receive data from the channel,
// and then the sender can hand it over directly to the receiver.
// So, channel send/receive operations block until the other side is ready
// A send operation on a channel (and the goroutine or function that contains it) blocks until a receiver is available for the same channel ch.
// A receive operation for a channel blocks (and the goroutine or function that contains it) until a sender is available for the same channel.

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
	close(ch)
}

func getData(ch chan string) {
	for input := range ch {
		fmt.Printf("%s ", input)
	}
	fmt.Println()
}

func Helper() {
	fmt.Println("From Channel Package!!!")
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
	pipeLineConcurrencyPattern()
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func square(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range ch {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func pipeLineConcurrencyPattern() {
	nums := []int{2, 4, 6, 8, 9}
	stage1 := sliceToChannel(nums)
	stage2 := square(stage1)
	for v := range stage2 {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

// buffered channel | Asynchronus | Non-blocking
// Sending to a buffered channel will not block unless the buffer is full (the capacity is completely used),
// and reading from a buffered channel will not block unless the buffer is empty.
// If the capacity is greater than 0, the channel is asynchronous
