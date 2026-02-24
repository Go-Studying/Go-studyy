package main

import "fmt"

func main() {
	fmt.Println("=== 버퍼 채널 ===")
	buffer()

	fmt.Println("=== 언버퍼 채널 ===")
	buffered()

	fmt.Println("=== 단방향 채널 ===")
	one_way()

	fmt.Println("=== 채널 세마포어 ===")
	channelSemaphore()
}