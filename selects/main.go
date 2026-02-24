package main

import "fmt"

func main() {
	fmt.Println("=== 다중 채널 Select ===")
	multiple()

	fmt.Println("=== 대기 없이 다른 코드 실행 ===")
	nonblock()

	fmt.Println("=== 타임아웃으로 제어하기 ===")
	timeout()

	fmt.Println("=== 동시성 도구 활용 ===")
	mixup()

	fmt.Println("=== mutex+select ===")
	mixup_with_mutex()
}
