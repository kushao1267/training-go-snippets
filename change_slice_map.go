package main

import (
	"fmt"
)

/*引用类型*/

func changeSlice(nums []int) {
	fmt.Printf("slice, before change: %p len: %d , cap: %d\n", nums, len(nums), cap(nums))
	nums = append(nums, 10) // cap为1，append后导致扩容创建新slice, 指针变化, 不影响原slice;
	fmt.Printf("slice, after change: %p len: %d , cap: %d\n", nums, len(nums), cap(nums))
	// nums[0] = 10 // 引用类型，所以直接更新值。
}

func changeMap(fakeMap map[int]int) {
	fmt.Printf("map, to change: %p \n", fakeMap)
	fakeMap[10] = 10
}

func main() {
	slice := []int{1}
	fmt.Printf("slice, before change: %v, pointer: %p , type: %T\n", slice, slice, slice)
	changeSlice(slice)
	fmt.Printf("slice, after change: %v, pointer: %p \n", slice, slice)
	fmt.Println("-----------------------------------------------------------")

	fakeMap := make(map[int]int)
	fakeMap[1] = 1
	// fakeMap := map[int]int{1: 1}
	fmt.Printf("map, before change: %v, pointer: %p , type: %T\n", fakeMap, fakeMap, fakeMap)
	changeMap(fakeMap)
	fmt.Printf("map, after change: %v, pointer: %p \n", fakeMap, fakeMap)
	fmt.Println("-----------------------------------------------------------")

	arr := [10]int{0}
	fmt.Printf("array, before change: %v, pointer: %p \n", arr, arr)
	arr[0] = 1
	fmt.Printf("array, after change: %v, pointer: %p \n", arr, arr)

}
