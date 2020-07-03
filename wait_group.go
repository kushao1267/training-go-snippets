package main

import (
	"fmt"
	"sync"
	"time"
)

func doWork(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	realSleepTime := time.Duration(n % 2)
	time.Sleep(realSleepTime * time.Second)
	fmt.Println("Do work: ", n)
}

func main() {
	// wg Add, Done, Wait三个主要的API
	var wg sync.WaitGroup
	N := 10

	for i := 0; i < N; i++ {
		wg.Add(1)
		go doWork(i, &wg) // wg是struct值类型，所以必须要传指针，才能修改其值，否则将会一直Wait下去，造成死锁!!
	}
	wg.Wait()
	fmt.Println("All done!")
}
