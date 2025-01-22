package main

import (
	"fmt"
)

func modifyArr(nums [3]int) {
	nums[0] = 2
	nums[1] = 1
	nums[2] = 0
}

func main() {
	a := [3]int{1, 2, 3}

	modifyArr(a)

	fmt.Println(a)
}
func SafeWrite(nums [5]int, i, val int) [5]int {
	if i >= len(nums) || i < 0 {
		return nums
	}
	nums[i] = val
	return nums
}
