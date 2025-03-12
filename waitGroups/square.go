package waitgroups

import (
	"fmt"
	"sync"
)

func square(index int, value int, res []int, wg *sync.WaitGroup) {
	defer wg.Done()
	res[index] = value * value

}

func SquareHelperViaWaitGroup() {
	arr := []int{1, 2, 3, 4, 5, 6}
	res := make([]int, len(arr))

	var wg sync.WaitGroup

	for i, v := range arr {
		wg.Add(1)
		go square(i, v, res, &wg)
	}

	wg.Wait()

	fmt.Println(res)

}
