package main

import "fmt"

func main() {
	const x string = "Hello World"
	fmt.Println(x)

	// 상수는 변할 수 없는 수이기 때문에 변수처럼 사용 불가
	// x = "Some other string"
}
