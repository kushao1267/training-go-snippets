package main

import (
	"fmt"
)

func main() {

	a := "abcdef"

	for i := range a {
		fmt.Printf("%c\n", a[i])
	}

	aBytes := []byte(a)
	aBytes[0] = 'w' // char赋值
	fmt.Printf("%c", string(aBytes))
}
