package main

import (
	"fmt"
)

func main() {

	a := make(chan struct{}, 5)
	a <- struct{}{}
	a <- struct{}{}
	fmt.Println(len(a))
}
