package main

import "fmt"
import "time"

func fibo(n uint, nmap []uint) uint {
	if n <= 1 {
		return 1
	}

	if nmap[n] == 0 {
		nmap[n] = fibo(n-1, nmap) + fibo(n-2, nmap)
	}

	return nmap[n]
}

func main() {
	max := 80
	jobCount := 3
	nmap := make([]uint, max)
	nmap[1] = 1
	nmap[0] = 1
	job := make(chan uint)
	done := make(chan bool, max)

	start := time.Now()
	// Spawns the jobs
	for j := 0; j < jobCount; j++ {
		go func(jNum int, iJob <-chan uint, umap []uint, done chan<- bool) {
			for n := range iJob {
				fmt.Printf("Job: %v running for %v\n", jNum, n)
				fibo(n, umap)
				done <- true
			}
		}(j, job, nmap, done)
	}

	for n := 0; n < max; n++ {
		job <- uint(n)
	}

	for i := 0; i < max; i++ {
		select {
		case <-done:
			fmt.Println("done")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout!!!!")
			i = max
		}
	}

	end := time.Now()

	diff := end.Sub(start)
	for j := 0; j < max; j++ {
		fmt.Println(j, nmap[uint(j)])
	}

	fmt.Println("Took", diff.Seconds(), "seconds")
}
