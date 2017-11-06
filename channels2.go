package main

import (
    "fmt"
    "time"
    )

func main(){
    msg := make(chan string)
    done := make(chan bool)

    fmt.Println("Starting")
    go func (messages <-chan string, done chan<- bool){
        for {
            select {
                case msg, more := <-messages:
                    if !more{
                        fmt.Println("Exiting worker")
                        done <- true
                        return
                    }

                    fmt.Println(msg)
                case <- time.After(time.Second * 2):
                    fmt.Println("Timeout")
            }
        }
    }(msg, done)


    for i := 0; i < 10; i++{
        msg <- fmt.Sprint("hello ", i)
        time.Sleep(time.Second * 2)
    }

    close(msg)
    <-done
}

