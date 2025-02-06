package main

import "fmt"

func main() {
	fmt.Println(GenerateSelfStory("Vlad", 25, 10.00000025))
}

func GenerateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%2.2f in my wallet right now.", name, age, money)
}
