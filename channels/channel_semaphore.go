package main

import (
	"fmt"
	"sync"
	"time"
)

func channelSemaphore() {
	sem := make(chan struct{}, 3) // 버퍼 크기 = 최대 동시 실행 수
	var wg sync.WaitGroup

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			sem <- struct{}{}        // acquire
			defer func() { <-sem }() // release

			fmt.Printf("Worker %d is starting\n", id)
			time.Sleep(2 * time.Second)
			fmt.Printf("Worker %d is done\n", id)
		}(i)
	}

	wg.Wait()
}