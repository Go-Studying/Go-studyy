package main

import (
	"fmt"
	"sync"
)

var (
	count int
)

func increment() {
	count++
}

func RaceCondition() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait() // 1000개의 고루틴이 끝날 때까지 대기
	fmt.Println("Count: ", count)
}
