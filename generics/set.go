package generics

// Sum() 함수는 제네릭 타입 T의 슬라이스의 합계에 대한 제약 조건을 정의
// T는 정수 or 부동 소수점 숫자
func Sum[T int | float64](numbers []T) T {
	var total T
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// 정수와 부동소수점 숫자의 슬라이스에 대해 Sum() 적용
}
