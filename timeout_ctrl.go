package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func(ch chan string) {
		time.Sleep(3 * time.Second)
		ch <- "Hello world"
	}(ch)

	select {
	case result := <-ch:
		fmt.Println("Task result is: ", result)
	case <-time.After(2 * time.Second):
		fmt.Println("The task is time out.")
		// default: // 立即判断所有channel，如果没有立即进入default
		// 	fmt.Println("Jump in default")
	}
}
