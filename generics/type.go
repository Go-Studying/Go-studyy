package generics

import "fmt"

// 제네릭 타입 정의
type Pair[T any] struct {
	First  T
	Second T
}

func main() {
	// 정수형 Pair
	intPair := Pair[int]{1, 2}
	fmt.Println("Integer Pair:", intPair)

	// 문자열 Pair
	strPair := Pair[string]{"hello", "world"}
	fmt.Println("String Pair:", strPair)
}
