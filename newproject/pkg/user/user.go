package user

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
}

func (u User) Greet() {
	fmt.Println("Добро пожаловать, ", u.Name)
}

func SayHello(name string) {
	fmt.Println("привет, ", name)
}
func Multiply(a, b int) int {
	return a * b
}
func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}
