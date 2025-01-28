package main

import (
	"fmt"
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
	resSlice := make([]int64, 0, 5)
	for i := 0; i < len(userIDs); i++ {
		resSlice = append(resSlice, userIDs[i])
		for j := 1; j < len(userIDs); j++ {
			if resSlice[i] == userIDs[j] {
				userIDs = append(userIDs[:j], userIDs[j+1:]...) //a = append(a[:1], a[2:]...)
				j--
			}
		}
	}
	return resSlice
}
