package main

import "fmt"

func main() {
	//fmt.Println(byte('a'))
	// fmt.Println(ShiftASCII("abc", 0))
	// fmt.Println(ShiftASCII("abc1", 1))
	// fmt.Println(ShiftASCII("bcd2", -1))
	// fmt.Println(ShiftASCII("hi", 10))
	// fmt.Println(ShiftASCII("abc", 256))
	// fmt.Println(ShiftASCII("abc", -511))
	fmt.Println("Hello helix")
	PrintSum(28, 35)
	PrintSum(45, 28)
	PrintSum(83, 324)
}

func PrintSum(n1 int, n2 int) {
	fmt.Println(n1 + n2)
}
func ShiftASCII(s string, step int) string {
	var rs string
	for _, b := range s {
		rs += string(byte(b) + byte(step))
	}
	return rs
}

// func NextASCII(b byte) byte {

// 	if (b + 1) > 255 {
// 		return b
// 	} else {
// 		return b + 1
// 	}

// }
// func PrevASCII(b byte) byte {
// 	if (b - 1) < 1 {
// 		return b
// 	} else {
// 		return b - 1
// 	}
// }

// Реализуйте функцию ShiftASCII(s string, step int) string, которая принимает на вход состоящую из ASCII символов строку s и возвращает новую строку, где каждый символ из входящей строки сдвинут вперед на число step.

// Например:

// ShiftASCII("abc", 0) // "abc"
// ShiftASCII("abc1", 1) // "bcd2"
// ShiftASCII("bcd2", -1) // "abc1"
// ShiftASCII("hi", 10) // "rs"
