package main

import (
	"fmt"
	"time"
)

func main() {
	var count1 *int //nil

	count2 := new(int) //adress
	countTemp := 5
	count3 := &countTemp //adress

	t := &time.Timer{} //C:(<-chan time.Time)(nil), initTimer:false
	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("time: %#v\n", *t)
		//fmt.Printf("time: %#v\n", t.String())
	}

}
