package main

import "fmt"

type Sayer interface {
	Say()
}
type Dog struct {
	name string
}

func (d *Dog) Say() {
	fmt.Println("wof wof")
}
func SaySmth(s Sayer) {
	s.Say()
}

func main() {
	fmt.Println("ks")
	SaySmth(&Dog{name: "Goofy"})
}

type OverflowErr struct {
    msg string
}

func (e error) (OverflowErr, error) {
    return e.msg, 
}