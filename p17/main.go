package main

import "fmt"

func main() {
	fmt.Println(MergeNumberLists([]int{1, 2}, []int{3}, []int{4}))
	fmt.Println(MergeNumberLists([]int{1, 2, 3}, []int{4}, []int{5, 6, 7, 8}, []int{9, 10}))
	fmt.Println(MergeNumberLists([]int{}))
	fmt.Println(MergeNumberLists(nil, nil))
	fmt.Println(MergeNumberLists([]int{10, 20}, nil, []int{50, 60}))
}

func MergeNumberLists(numberLists ...[]int) []int {
	var res = []int{}
	for _, s := range numberLists {
		if s == nil {
			continue
		}
		res = append(res, s...)
	}
	return res
}
