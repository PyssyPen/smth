package main

import (
	"fmt"
)

func MostPopularWord(words []string) string {
	wordsCount := make(map[string]int, 0)
	mostPopWord := ""
	max := 0

	for _, word := range words {
		fmt.Print(wordsCount[word])
		fmt.Println(word)
		wordsCount[word]++
		if wordsCount[word] > max {
			max = wordsCount[word]
		}
	}

	for _, word := range words {
		if wordsCount[word] == max {
			mostPopWord = word
			break
		}
	}

	return mostPopWord
}

func main() {
	l := 0
	fmt.Print("Длина слайса: ")
	fmt.Scan(&l)
	sl := make([]string, 0, l)
	for i := 0; i < l; i++ {
		fmt.Print("Введите ", i+1, " слово: ")
		slovo := ""
		fmt.Scanln(&slovo)
		sl = append(sl, slovo)
	}
	fmt.Println("Полученный слайс:", sl)
	r := MostPopularWord(sl)
	fmt.Println("Самое популярное слово:", r)
}
