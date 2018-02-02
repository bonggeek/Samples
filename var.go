package main

import "fmt"

func main() {
	var a string = "intial"
	var b, c string = "foo", "bar"

	d := "la"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	const pi = 22e0 / 7
	fmt.Println(float64(pi))
	fmt.Println((pi))
}
