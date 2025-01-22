package main

import "fmt"

const (
	zero = 0
	one
	two
	three = iota
	four
	five
)

func main() {
	fmt.Println(zero, one, two, three, four, five)
}
