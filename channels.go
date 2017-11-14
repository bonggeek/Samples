package main

import (
	"fmt"
	"time"
)

func f(text string, in <-chan bool, out chan<- bool) {
	for true {
		<-in
		fmt.Println(text)
		time.Sleep(time.Second)
		out <- true
	}
}

func main() {

	chan1 := make(chan bool)
	chan2 := make(chan bool)
	go f("ping", chan1, chan2)
	go f("pong", chan2, chan1)

	chan1 <- true

	fmt.Println("Channel size", cap(chan1))
	var input string
	fmt.Scanln(&input)
}
