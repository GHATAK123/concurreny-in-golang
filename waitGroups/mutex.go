package waitgroups

import (
	"fmt"
	"sync"
	"time"
)

var lock sync.Mutex

func ProcessValue(data int) int {
	time.Sleep(time.Second * 2)
	return data * 2
}

func ProcessData(wg *sync.WaitGroup, res *[]int, data int) {

	defer wg.Done()
	val := ProcessValue(data)
	lock.Lock()
	*res = append(*res, val)
	lock.Unlock()
}

func DemoHelper() {

	start := time.Now()
	var wg sync.WaitGroup

	given := []int{1, 2, 3, 4, 5, 6}
	result := []int{}

	for _, data := range given {
		wg.Add(1)
		go ProcessData(&wg, &result, data)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
	fmt.Println(result)

}
