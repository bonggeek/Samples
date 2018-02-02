package main

import "fmt"

func cl(i int) func(int) int {
	return func(j int) int {
		i++
		return i * j
	}
}

func main() {
	var f1 func(int) int
	f1 = cl(10)
	fmt.Println(f1(3))
	fmt.Println(f1(3))

	f2 := cl(41)
	fmt.Println(f2(3))
	fmt.Println(f2(3))
}
