package main

import "fmt"

var y string = "Hello World"

func main() {
	var x string = "Hello World"
	fmt.Println(x)

	fmt.Println(y)
	f()
}

func f() {
	fmt.Println(y)
}
