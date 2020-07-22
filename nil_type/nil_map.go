package main

import (
	"fmt"
)

// func main() {
// 	// nil map查值，返回对应类型的零值。
// 	var intMap map[int]int
// 	v, ok := intMap[1]
// 	fmt.Println(v, ok)

// 	// delete key
// 	// intMap = map[int]int{1: 1}
// 	// delete(intMap, 1)

// 	// nil map设值报错
// 	intMap[1] = 1
// }
func main() {
	// var s = make([]string, 10)   // 初始化10个空字符串元素
	var s []string // 初始化为空[]
	s = append(s, "aaa")

	for i, v := range s {
		fmt.Printf("s[%d]=[%s]\n", i, v)
	}
}
