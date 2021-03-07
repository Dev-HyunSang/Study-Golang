package main

import "fmt"

func main() {
	var a string = "Hello World"
	var b string
	b = "Hello World"
	fmt.Println(a)
	fmt.Println(b)

	var c string
	c = "first"
	fmt.Println(c)
	c = "second"
	fmt.Println(c)

	var d string
	d = "first"
	fmt.Println(d)
	d = d + " second"
	fmt.Println(d)

	var x string = "Hello"
	var y string = "World"
	fmt.Println(x == y)

	x = "Hello"
	y = "Hello"
	fmt.Println(x == y)

	z := "Hello World"
	fmt.Println(z)

	var e = "Hello World"
	fmt.Println(e)

	f := 5
	fmt.Println(f)
}
