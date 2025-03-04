package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const numRequests = 10_000

var count int32

func networkRequest(wg *sync.WaitGroup) {
	defer wg.Done()
	// defer wg.Done() // - panic: negative WaitGroup counter

	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса.
	atomic.AddInt32(&count, 1)
}

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(numRequests)
	timeStart := time.Now()

	for i := 0; i < numRequests; i++ {
		go networkRequest(&wg)
	}

	wg.Wait()
	fmt.Println(count) // 10000, 10s

	timeEnd := time.Now()
	diff := timeEnd.Sub(timeStart)

	fmt.Printf("Duration: %d \n", diff.Milliseconds())
	// Сколько по времени ?
	// 1. 11707 - 12750 ms
	// 2. 5 - 10 ms

	// 10_000			: 5 - 10
	// 100_000			: 36
	// 1_000_000		: 262
	// 10_000_000		: 2_523
	// 100_000_000		: 24_707
	// 1_000_000_000	: 245_748
}
