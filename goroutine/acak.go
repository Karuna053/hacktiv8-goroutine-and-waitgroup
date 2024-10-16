package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

func GoroutineAcak() {
	fmt.Println("Goroutine Acak")

	var wg sync.WaitGroup

	count := 4

	// async function in here with goroutine
	for i := 1; i <= count; i++ {
		wg.Add(1)
		// ini proses goroutine
		go bisaProcess(i, &wg)
	}

	for i := 1; i <= count; i++ {
		wg.Add(1)
		// ini proses goroutine
		go cobaProcess(i, &wg)
	}

	// printing banyaknya goroutine yang jalan
	fmt.Println(runtime.NumGoroutine())

	wg.Wait() // Waits for all instance of wait group to be finished.
}

func bisaProcess(index int, wg *sync.WaitGroup) {
	data := []interface{}{"bisa1", "bisa2", "bisa3"}
	fmt.Println(data, index)
	wg.Done() // informs that this instance of wait in the wait group is done,
}

func cobaProcess(index int, wg *sync.WaitGroup) {
	data := []interface{}{"coba1", "coba2", "coba3"}
	fmt.Println(data, index)
	wg.Done() // informs that this instance of wait in the wait group is done,
}
