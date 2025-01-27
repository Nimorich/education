package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println(UniqueSortedUserIDs([]int64{}))
	fmt.Println(UniqueSortedUserIDs([]int64{10}))
	fmt.Println(UniqueSortedUserIDs([]int64{55, 55}))
	fmt.Println(UniqueSortedUserIDs([]int64{55, 55, 33, 22}))
	fmt.Println(UniqueSortedUserIDs([]int64{55, 2, 88, 33, 2, 2, 55, 103, 33, 88}))
}

func UniqueSortedUserIDs(userUDs []int64) []int64 {
	sort.Slice(userUDs, func(i, j int) bool {
		return userUDs[i] < userUDs[j]
	})
	return slices.Compact(userUDs)
}

// func UniqueSortedUserIDs(userIDs []int64) []int64 {
// 	if len(userIDs) < 2 {
// 		return userIDs
// 	}

// 	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
// 	uniqPointer := 0
// 	for i := 1; i < len(userIDs); i++ {
// 		if userIDs[uniqPointer] != userIDs[i] {
// 			uniqPointer++
// 			userIDs[uniqPointer] = userIDs[i]
// 		}
// 	}

// 	return userIDs[:uniqPointer+1]
