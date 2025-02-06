package main

import (
	// "fmt"
	"strings"
	"unicode"
)

func main() {
	// sb := strings.Builder{}
	// sb.WriteString("hello")
	// sb.WriteString(",")
	// sb.WriteString(" ")
	// sb.WriteString("go")

	// fmt.Println(sb.String())

}
func LatinLetters(s string) string {
	sb := strings.Builder{}
	for _, r := range s {
		if unicode.Is(unicode.Latin, r) {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
