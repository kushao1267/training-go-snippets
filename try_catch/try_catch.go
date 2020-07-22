package main

import "fmt"

func try(f func()) { // 包装可能会发生panic的函数
	defer func() {
		if err := recover(); err != nil { // 注意: recover只能处理当前Goroutine!!
			fmt.Println("Panic Occured.", err)
		}
	}()
	f()
}

func work() {
	panic("error 1")
}

func main() {
	// work()
	try(work)
}
