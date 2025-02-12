package main

import "fmt"

type Counter struct {
	Value int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	c := Counter{}
	c.Inc(0)
	c.Inc(0)
	c.Inc(40)
	fmt.Println("Inc40 %d", c.Value)

	c.Dec(0)
	c.Dec(30)
	// 11
	fmt.Println("Dec30 %d", c.Value)

	c.Dec(100)
	// 0
	fmt.Println("Dec100 %d", c.Value)
}

func (c *Counter) Inc(delta int) {
	if delta == 0 {
		delta++
	}
	c.Value += delta
}
func (c *Counter) Dec(delta int) {
	if delta == 0 {
		delta++
	}
	if c.Value-delta < 0 {
		c.Value = 0
	} else {
		c.Value -= delta
	}
}
