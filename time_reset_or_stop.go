package main

import (
	"fmt"
	"time"
)

func main() {
	// 计时开始
	start := time.Now()
	timeArrival := make(chan struct{})
	// 2s后执行func
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("Hello world")
		close(timeArrival) // 与timeArrival <- struct{}{}效果相同
	})
	time.Sleep(1 * time.Second)

	// 1.stop timer
	// timer.Stop()
	// fmt.Println("Stop task, Elaspe:", time.Since(start))

	// 2.reset timer
	timer.Reset(4 * time.Second)
	<-timeArrival
	fmt.Println("Reset task, Elaspe:", time.Since(start))
	// Elaspe: 5.001394516s 因为sleep 1秒后重置为4s，所以需要5秒
}
