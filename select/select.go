package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		检查每个case代码块
		如果任意一个case代码块准备好发送或接收，执行对应内容
		如果多余一个case代码块准备好发送或接收，`随机`选取一个并执行对应内容
		如果任何一个case代码块都没有准备好，等待
		如果有default代码块，并且没有任何case代码块准备好，执行default代码块对应内容
	*/
	a := make(chan int, 1)
	b := make(chan int, 1)
	c := make(chan int, 1)
	a <- 1
	b <- 1
	c <- 1

loop: // break loop方式跳出for循环
	for {
		select {
		case v := <-a:
			fmt.Println("a", v)
		case v := <-b:
			fmt.Println("b", v)
		case v := <-c:
			fmt.Println("c", v)
		default:
			fmt.Println("go out")
			break loop
		}
	}
	fmt.Println("Done")
	errors.Wrapf
}
