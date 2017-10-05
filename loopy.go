package main

import "fmt"

func main(){
    i := 1
    for i <= 10{
        fmt.Println(i)
        i = i + 1
    }

    var j int = 0

    for ; j <= 42; j++{
        fmt.Println(j)
    }
}
