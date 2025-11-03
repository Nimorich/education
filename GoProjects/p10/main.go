package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(UniqueUserIDs([]int64{}))
	fmt.Println(UniqueUserIDs([]int64{10}))
	fmt.Println(UniqueUserIDs([]int64{55, 55}))
	fmt.Println(UniqueUserIDs([]int64{55, 55, 33, 22}))
	fmt.Println(UniqueUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}))
}

// нужно пройтись по слайсу и удалить все похожие значения
func UniqueUserIDs(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}
	resSlice := make([]int64, 0, len(userIDs))
	resSlice = append(resSlice, userIDs[0])
	index := 0
	for i, v := range userIDs {
		if slices.Contains(resSlice, v) {
			continue
		} else {
			resSlice = append(resSlice, userIDs[i])
			index++
		}
	}

	return resSlice

}
