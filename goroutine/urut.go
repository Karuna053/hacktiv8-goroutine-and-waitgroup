package goroutine

import (
	"fmt"
	"sync"
)

func GoroutineUrut() {
	fmt.Println("Goroutine Urut")

	var wg sync.WaitGroup // Basic waitgroup function
	var locker sync.Mutex // This is literally to bikin antri no matter what

	count := 8

	for i := 1; i <= count; i++ {
		wg.Add(1)
		locker.Lock()

		if i%2 == 00 {
			go bisaProcessUrut(i, &wg, &locker)
		} else {
			go cobaProcessUrut(i, &wg, &locker)
		}
	}

	// fmt.Println(runtime.NumGoroutine())

	wg.Wait()
}

func bisaProcessUrut(index int, wg *sync.WaitGroup, locker *sync.Mutex) {
	data := []interface{}{"bisa1", "bisa2", "bisa3"}
	fmt.Println(data, index)
	locker.Unlock()
	wg.Done()
}

func cobaProcessUrut(index int, wg *sync.WaitGroup, locker *sync.Mutex) {
	data := []interface{}{"coba1", "coba2", "coba3"}
	fmt.Println(data, index)
	locker.Unlock()
	wg.Done()
}
