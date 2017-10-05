package main

import "fmt"

func cl(i int) func () int {
    return func () int {
        i++
        return i
    }
}

func main() {
    var f1 func() int
    f1 = cl(10);
    fmt.Println(f1())
    fmt.Println(f1())

    f2 := cl(41);
    fmt.Println(f2())
    fmt.Println(f2())
}
