package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Andy", Age: 18}
	fmt.Println("simple struct:", p)
	fmt.Printf("detailed struct: %+v\n", p)
	fmt.Printf("Golang struct: %#v\n", p)
}
