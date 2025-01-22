package main

import (
	"fmt"
	"strings"
)

/*
 */
func main() {
	fmt.Println(Validate(UserCreateRequest{
		FirstName: "",
		Age:       0,
	}))
	fmt.Println(Validate(UserCreateRequest{
		FirstName: "John",
		Age:       -5,
	}))
	fmt.Println(Validate(UserCreateRequest{
		FirstName: "Andy",
		Age:       0,
	}))
	fmt.Println(Validate(UserCreateRequest{
		FirstName: "Karl",
		Age:       151,
	}))
	fmt.Println(Validate(UserCreateRequest{
		FirstName: "",
		Age:       5,
	}))
	fmt.Println(Validate(UserCreateRequest{
		FirstName: " Henry",
		Age:       15,
	}))
	fmt.Println(Validate(UserCreateRequest{
		FirstName: "John Smith",
		Age:       15,
	}))
	fmt.Println(Validate(UserCreateRequest{
		FirstName: "John",
		Age:       150,
	}))
	fmt.Println(Validate(UserCreateRequest{
		FirstName: "Susan",
		Age:       30,
	}))
}

type UserCreateRequest struct {
	FirstName string
	Age       int
}

// BEGIN (write your solution here)
func Validate(req UserCreateRequest) string {
	if strings.Contains(req.FirstName, " ") || req.Age <= 0 {
		return "invalid request"
	} else if req.FirstName == "" || req.Age > 150 {
		return "invalid request"
	} else {
		return ""
	}
}
