package main

import (
	"fmt"
)

func UniqueUserIDs(userIDs []int64) []int64 {
	// пустая структура struct{} — это тип данных, который занимает 0 байт
	// используется, когда нужно проверять в мапе только наличие ключа
	processed := make(map[int64]struct{})

	uniqUserIDs := make([]int64, 0)
	for _, uid := range userIDs {

		_, ok := processed[uid]
		fmt.Print("uid ", uid)
		if ok {
			fmt.Println(" NAHUI")
			continue
		}
		uniqUserIDs = append(uniqUserIDs, uid)
		fmt.Println(" OK")
		processed[uid] = struct{}{}
	}
	return uniqUserIDs
}

func main() {
	l := 0
	fmt.Print("Длина слайса: ")
	fmt.Scanln(&l)
	sl := make([]int64, 0, l)
	for i := 0; i < l; i++ {
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
	unique := UniqueUserIDs(sl)
	fmt.Println("Измененный слайс:", unique)
}
