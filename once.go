package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once // 即使传入的函数不同，也只会执行第一个函数
	onceBody := func() {
		fmt.Println("Only once")
	}

	onceBody1 := func() {
		fmt.Println("Only once 1")
	}

	once.Do(onceBody)
	once.Do(onceBody1)
}
