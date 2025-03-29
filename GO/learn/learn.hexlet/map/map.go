package main

import (
	"fmt"
	"sort"
)

/*
Реализуйте функцию UniqueUserIDs(userIDs []int64) []int64, которая возвращает слайс,
состоящий из уникальных идентификаторов userIDs. Порядок слайса должен сохраниться.
*/
func uUniqueSortedUserIDs(userIDs []int64) []int64 {
	elementMap := make(map[int64]int64) // создается пустая мапа где ключ = числу

	for _, s := range userIDs {
		elementMap[s] = s
		fmt.Println(elementMap)
		// с каждой новой итерацией прщверяется, есть ли ключ, и если его нет, то добавляет его и его значение (ключ = значение)
	}

	v := make([]int64, 0, len(elementMap)) // создаеся пустой слайс
	for _, value := range elementMap {
		v = append(v, value)
		fmt.Println(v)
	}
	sort.Slice(v, func(i, j int) bool {
		fmt.Println(v)
		return v[i] < v[j]
	})
	return v
}

func UNmain() { //main
	l := 0
	fmt.Print("Длина слайса: ")
	fmt.Scanln(&l)
	sl := make([]int64, 0, l)

	for i := l - 1; i > -1; i-- {
		m := int64(i)
		if m%2 == 0 {
			sl = append(sl, m+1)
			sl = append(sl, (m*2)+1)
			sl = append(sl, 1)
		} else {
			sl = append(sl, m+1)
			sl = append(sl, m+2)
		}
	}

	fmt.Println("Полученный слайс:", sl)
	r := uUniqueSortedUserIDs(sl)
	fmt.Println("Измененный слайс:", r)
}

/*
Вывод программы:

Длина слайса: 6
Полученный слайс: [6 7 5 9 1 4 5 3 5 1 2 3 1 1 1]
map[6:6]
map[6:6 7:7]
map[5:5 6:6 7:7]
map[5:5 6:6 7:7 9:9]
map[1:1 5:5 6:6 7:7 9:9]
map[1:1 4:4 5:5 6:6 7:7 9:9]
map[1:1 4:4 5:5 6:6 7:7 9:9]
map[1:1 3:3 4:4 5:5 6:6 7:7 9:9]
map[1:1 3:3 4:4 5:5 6:6 7:7 9:9]
map[1:1 3:3 4:4 5:5 6:6 7:7 9:9]
map[1:1 2:2 3:3 4:4 5:5 6:6 7:7 9:9]
map[1:1 2:2 3:3 4:4 5:5 6:6 7:7 9:9]
map[1:1 2:2 3:3 4:4 5:5 6:6 7:7 9:9]
map[1:1 2:2 3:3 4:4 5:5 6:6 7:7 9:9]
map[1:1 2:2 3:3 4:4 5:5 6:6 7:7 9:9]
[1]
[1 4]
[1 4 3]
[1 4 3 2]
[1 4 3 2 6]
[1 4 3 2 6 7]
[1 4 3 2 6 7 5]
[1 4 3 2 6 7 5 9]
[1 4 3 2 6 7 5 9]
[1 4 3 2 6 7 5 9]
[1 3 4 2 6 7 5 9]
[1 3 4 2 6 7 5 9]
[1 3 2 4 6 7 5 9]
[1 2 3 4 6 7 5 9]
[1 2 3 4 6 7 5 9]
[1 2 3 4 6 7 5 9]
[1 2 3 4 6 7 5 9]
[1 2 3 4 6 5 7 9]
[1 2 3 4 5 6 7 9]
[1 2 3 4 5 6 7 9]
Измененный слайс: [1 2 3 4 5 6 7 9]
*/
