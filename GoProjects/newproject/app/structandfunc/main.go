package main

import (
	"fmt"
	"newproject/pkg/user"
)

func main() {
	user.SayHello("Алиса")

	result := user.Multiply(3, 4)
	fmt.Println("Результат умножения:", result)

	quotient, err := user.SafeDivide(10, 2)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Результат деления:", quotient)
	}

	user := user.User{Name: "Сосиса"}
	user.Greet()
}
