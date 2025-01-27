package main

import "fmt"

func main() {
	fmt.Println(IntsCopy([]int{}, 0))
	fmt.Println(IntsCopy([]int{1, 2}, 0))
	fmt.Println(IntsCopy([]int{1, 2}, -1))
	fmt.Println(IntsCopy([]int{1, 2}, -5))
	fmt.Println(IntsCopy([]int{1, 2, 3}, 2))
	fmt.Println(IntsCopy([]int{1, 2, 3, 4}, 4))
	fmt.Println(IntsCopy([]int{1, 2, 3, 4, 5}, 10))
}

/*
a.Equal([]int{}, pkg.IntsCopy([]int{}, 0))
a.Equal([]int{}, pkg.IntsCopy([]int{1, 2}, 0))
a.Equal([]int{}, pkg.IntsCopy([]int{1, 2}, -1))
a.Equal([]int{}, pkg.IntsCopy([]int{1, 2}, -5))
a.Equal([]int{1, 2}, pkg.IntsCopy([]int{1, 2, 3}, 2))
a.Equal([]int{1, 2, 3, 4}, pkg.IntsCopy([]int{1, 2, 3, 4}, 4))
a.Equal([]int{1, 2, 3, 4, 5}, pkg.IntsCopy([]int{1, 2, 3, 4, 5}, 10))
*/
func IntsCopy(src []int, maxLen int) []int {
	var res []int
	if maxLen >= 0 {
		if maxLen > len(src) {
			res = make([]int, len(src))
			copy(res, src)
			return res
		}
		res = make([]int, maxLen)
		copy(res, src)
		return res
	} else if maxLen <= 0 {
		res = make([]int, 0)
		return res
	}
	return res
}
