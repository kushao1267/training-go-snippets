package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// var ch chan int // nil chan leads to deadlock!
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	//get first result
	for i := 0; i < 3; i++ {
		fmt.Println("result:", <-ch)
	}

	//do other work
	time.Sleep(1 * time.Second)
}
