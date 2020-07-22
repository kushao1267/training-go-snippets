package main

import (
	"fmt"
)

/*引用类型*/

func changeSlice(nums []int) {
	fmt.Printf("slice, before change: value: %v, %p len: %d , cap: %d\n", nums, nums, len(nums), cap(nums))
	nums = append(nums, 10) // cap为1，append后导致扩容创建新slice, 指针变化, 不影响原slice;
	fmt.Printf("slice, after change: value: %v, %p len: %d , cap: %d\n", nums, nums, len(nums), cap(nums))
}

func changeMap(fakeMap map[int]int) {
	fmt.Printf("map, to change: %p \n", fakeMap)
	fakeMap[10] = 10
}

func main() {
	/*结论: slice传指针地址，在函数内append发生扩容将会创建新slice，指针的值将会变化，导致原slice的append失败*/
	fmt.Println("-------------------------change slice---------------------------")
	slice := []int{1}
	fmt.Printf("slice, before change: %v, pointer: %p \n", slice, slice)
	changeSlice(slice)
	fmt.Printf("slice, after change: %v, pointer: %p \n", slice, slice)

	fmt.Println("-------------------------change map---------------------------")
	fakeMap := make(map[int]int, 1)
	fakeMap[1] = 1
	// fakeMap := map[int]int{1: 1}
	fmt.Printf("map, before change: %v, pointer: %p, len: %d\n", fakeMap, fakeMap, len(fakeMap))
	changeMap(fakeMap)
	fmt.Printf("map, after change: %v, pointer: %p, len: %d\n", fakeMap, fakeMap, len(fakeMap))
}
