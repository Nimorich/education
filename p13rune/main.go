package main

import "fmt"

func main() {
	fmt.Println(IsASCII(" abc1"))
	fmt.Println(IsASCII("хекслет"))
	fmt.Println(IsASCII("hello \U0001F970"))
}

func IsASCII(s string) bool {
	for _, c := range s {
		if c > 255 {
			return false
		}
	}
	return true
}
