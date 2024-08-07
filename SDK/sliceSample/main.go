package main

import "fmt"

func main() {
	nums := make([]int, 5, 7)
	numsCut := nums[1:4]
	fmt.Println(numsCut, "cap: ", cap(numsCut), "len: ", len(numsCut))
	
	sliceOps(numsCut)
	fmt.Println(numsCut, "cap: ", cap(numsCut), "len: ", len(numsCut))
	fmt.Println(nums, "cap: ", cap(nums), "len: ", len(nums))
}

func sliceOps(nums []int) {
	// append 大于切片容量时，会重新分配数组，导致数组地址发生变化，因此不会影响原切片
	nums = append(nums, 1, 2, 3)
	nums[0] = 5
	fmt.Println(nums, "cap: ", cap(nums), "len: ", len(nums))
}
