package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
    fmt.Println("jsa")
	//fmt.Println("Hello")
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range arguments[1:] {
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}
		invalid = append(invalid, k)
	}
	fmt.Printf("#read: %d #ints: %d #floats: %d \n", total, nInts, nFloats)
	fmt.Println("Too much invalid input: ", invalid)
}
