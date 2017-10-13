package main

import "fmt"
import "time"

func fibo(n uint) uint{
    if n <= 1 {
        return 1
    }

    return fibo(n - 1) + fibo(n - 2)
}

func main() {
    start := time.Now()
    for n := 1; n < 43; n++{
        fmt.Println(n, fibo(uint(n)))
    }

    end := time.Now()

    diff := end.Sub(start)
    fmt.Println("Took", diff.Seconds(), "seconds")
}
