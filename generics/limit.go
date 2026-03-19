package generics

import "fmt"

// 제네릭 함수 정의
func PrintLength[T any](value T) {
	// 다음 코드는 T가 string 또는 [] int 일때만 동작
	// 그러나 제네릭에는 이러한 타입 검사를 컴파일 타임에 보장할 수 없다.
	fmt.Println(len(value))
}

func main() {
	PrintLength("hello")        // 작동
	PrintLength([]int{1, 2, 3}) // 작동
	// PrintLength(42) // 컴파일 오류: int 형은 len() 함수를 지원하지 않는다.
}
