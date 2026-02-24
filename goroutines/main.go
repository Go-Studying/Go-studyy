package main

import "fmt"

func main() {
	fmt.Println("=== 고루틴 흐름 ===")
	flow()

	fmt.Println("=== 레이스 컨디션 ===")
	RaceCondition()

	fmt.Println("=== WaitGroup ===")
	waitgroup()
}