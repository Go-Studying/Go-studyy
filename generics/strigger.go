package generics

import "fmt"

// Strigger 인터페이스 정의
type Strigger interface {
	String() string
}

// 제네릭 함수 정의
func PrintString[T Strigger](value T) {
	fmt.Println(value.String())
}

// String 인터페이스를 구현하는 타입 정의
type Person struct {
	Name string
}

// String() 메서드 구현
func (p Person) String() string {
	return p.Name
}

func main() {
	// Person 타입을 사용하는 제네릭 함수
	p := Person{Name: "John Doe"}
	PrintString(p)
}
