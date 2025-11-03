package main

import (
	"fmt"
)

func main() {
	p := &Parent{
		Name: "Harry",
		Children: []Child{
			{
				Name: "Andy",
				Age:  18,
			},
		},
	}
	fmt.Println(p)
	cp := CopyParent(nil)
	fmt.Println(cp)
	cp = CopyParent(p)
	cp.Children[0] = Child{}
	fmt.Println(cp)
	fmt.Println(p)
}

type Parent struct {
	Name     string
	Children []Child
}
type Child struct {
	Name string
	Age  int
}

// BEGIN (write your solution here)
func CopyParent(p *Parent) Parent {
	if p == nil {
		rn := new(Parent)
		return *rn
	}
	rp := *p
	rp.Children = make([]Child, len(p.Children))
	copy(rp.Children, p.Children)
	return rp
}
