package main

import "fmt"

type testStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (t *testStruct) Shoot() bool {
	if !t.On {
		return false
	}
	if t.Ammo > 0 {
		t.Ammo--
		return true
	} else {
		return false
	}
}

// RideBike это метод, который изменяет testStruct
func (t *testStruct) RideBike() bool {
	if !t.On {
		return false
	}
	if t.Power > 0 {
		t.Power--
		return true
	} else {
		return false
	}
}
func main() {
	t := testStruct{true, 10, 10}
	testStruct := &t
	fmt.Println(testStruct.Shoot(), " Количество Ammo = ", testStruct.Ammo)
	fmt.Println(testStruct.RideBike(), " Количество Power = ", testStruct.Power)
	var notes [3]string = [3]string{
		"do",
		"re",
		"mi",
	}
	fmt.Println(notes)
}
