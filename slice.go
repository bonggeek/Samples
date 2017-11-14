package main

import "fmt"

func main() {
	s1 := []byte{0, 1, 2, 3, 4, 5}
	s2 := s1[2:4]
	s3 := make([]byte, len(s1))
	copy(s3, s1)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	s2[0] = 10
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
