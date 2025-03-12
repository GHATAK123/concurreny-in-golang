package channels

import (
	"fmt"
	"sync"
)

type Result struct {
	index, value int
}

func SquareElement(index int, num int, result chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	result <- Result{index, num * num}
}

func SquareHelper() {
	arr := []int{1, 2, 3, 4, 5, 6}
	ch := make(chan Result, len(arr))

	var wg sync.WaitGroup

	for i, num := range arr {
		wg.Add(1)
		go SquareElement(i, num, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	squareSlice := make([]int, len(arr))

	for v := range ch {
		squareSlice[v.index] = v.value
	}

	fmt.Println(squareSlice)

}
