package main

import "fmt"

type student struct {
    class int
    name string
    grade byte
}

func main() {
    students := [...]student {
        student{class: 1, name: "abhinaba", grade: 'a'},
        student{class: 1, name: "somtapa", grade: 'b'},
        student{class: 1, name: "pori", grade: 'c'},
        student{class: 1, name: "nogen", grade: 'c'},
   

    var s *student = &students[0]
    s.class = 4

    fmt.Println(students)
}
