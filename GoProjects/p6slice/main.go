package main

import (
	"fmt"
)

func modifySlice(nums []int) {
	nums = append(nums, 4)
	nums[0] = 2
	nums[1] = 1
	nums[2] = 0
}

func main() {
	// a := []int{1, 2, 3}
	// modifySlice(a)
	// fmt.Println(a)
	fmt.Println(RemoveFirstElement([]int{1, 2, 3, 4}))
	fmt.Println(RemoveFirstElement([]int{}))
	fmt.Println(RemoveFirstElement([]int{-4, -3, -2}))
	fmt.Println(RemoveFirstElement([]int{3}))
}
func RemoveFirstElement(slice []int) []int {
	if len(slice) > 0 {
		var srez []int = make([]int, len(slice)-1)
		for i := 0; i < len(srez); i++ {
			srez[i] = slice[i+1]
		}
		return srez
	} else {
		emptySlice := []int{}
		return emptySlice
	}

}
