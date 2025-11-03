package main

import "fmt"

type Grades []float64

func (g *Grades) add(grade float64) {
	*g = append(*g, grade)
}

func main() {
	grades := Grades{3.7, 4.2, 5.0}
	(&grades).add(4.8)
	fmt.Println(grades) // [3.7 4.2 5 4.8]
}
