package main

import "fmt"

func main() {
	var i int
	fmt.Scanf("%d", &i)

	if i == 0 {
		fmt.Println("영")
	} else if i == 1 {
		fmt.Println("하나")
	} else if i == 2 {
		fmt.Println("둘")
	} else if i == 3 {
		fmt.Println("삼")
	} else if i == 4 {
		fmt.Println("넷")
	} else if i == 5 {
		fmt.Println("다섯")
	}

	switch i {
	case 0:
		fmt.Println("영")
	case 1:
		fmt.Println("일")
	case 2:
		fmt.Println("둘")
	case 3:
		fmt.Println("삼")
	case 4:
		fmt.Println("넷")
	case 5:
		fmt.Println("다섯")
	}
}
