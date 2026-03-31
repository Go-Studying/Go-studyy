package main

import (
	"fmt"
	"net/http"
	"time"
)

func optimizedCommunication() {
	ch := make(chan int, 100)

	for i := 0; i < 100; i++ {
		go func(n int) {
			time.Sleep(1 * time.Millisecond) // 인위적인 지연
			ch <- n
		}(i)
	}

	for i := 0; i < 100; i++ {
		fmt.Println(<-ch) // 수신된 값 출력
	}
}

func run1() {
	go optimizedCommunication()
	http.ListenAndServe("localhost:6060", nil)
}
