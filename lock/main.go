package main

import "fmt"

func main() {
	fmt.Println("=== Mutex Count ===")
	countWithMutex()

	fmt.Println("=== Semaphore ===")
	semaphoreExample()

	fmt.Println("=== WaitGroup ===")
	waitGroup()
}