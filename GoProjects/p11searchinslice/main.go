package main

import "fmt"

func main() {
	fmt.Println(MostPopularWord([]string{"hello", "world", "hello"}))
	fmt.Println(MostPopularWord([]string{"hello", "world", "hello", "world", "world"}))
	fmt.Println(MostPopularWord([]string{"one", "two", "three", "four", "five"}))
	fmt.Println(MostPopularWord([]string{"a", "b", "c", "c", "d", "e", "e", "d"}))
	fmt.Println(MostPopularWord([]string{"a", "c", "c", "a"}))
}
func MostPopularWord(words []string) string {

	min := 0
	max := 0

	maxV := words[0]
	for i, word := range words { //итерация первого значения
		minV := words[i]
		for w2 := i + 1; w2 < len(words); w2++ { //итерация второго значения
			if words[w2] == word { //если равны
				min++
				minV = words[w2] //добавляем в мин Значение words[w2]
				if min > max {   //если мин больше макс
					maxV = minV //добавляем в макс Значение minV
					max = min
				}
			}
		}
		min = 0
	}
	return maxV
}

// func MostPopularWord(words []string) string {
// 	wordsCount := make(map[string]int, 0)
// 	mostPopWord := ""
// 	max := 0
// 	for _, word := range words {
// 		wordsCount[word]++
// 		if wordsCount[word] > max {
// 			max = wordsCount[word]
// 			mostPopWord = word
// 		}
// 	}
// 	return mostPopWord
// }
