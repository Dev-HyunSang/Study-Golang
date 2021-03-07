package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			// 나머지가 0이면 짝수
			fmt.Println(i, "짝수")
		} else {
			// 나머지가 0이 아닌 경우 홀수
			fmt.Println(i, "홀수")
		}
	}
}
