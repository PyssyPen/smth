package main

import (
	"fmt"
	"sort"
)

//  Реализуйте функцию UniqueSortedUserIDs(userIDs []int64) []int64,
//  которая возвращает отсортированный слайс, состоящий из уникальных идентификаторов userIDs.

/* Решениее с гитхаба
func UUniqueSortedUserIDs(userIDs []int64) []int64 {
	elementMap := make(map[int64]int64)
	for _, s := range userIDs {
		elementMap[s] = s
	}
	v := make([]int64, 0, len(elementMap))

	for _, value := range elementMap {
		v = append(v, value)
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i] < v[j]
	})
	return v
}



ЭТО ДОЛЖНО РАБОТАТЬ НО ТЕСТЫ ВЫДАЮТ ОШИБКУ

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}
	lenght := len(userIDs)
	up := int64(0)
	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {

			sort.Slice(userIDs, func(l, m int) bool { return userIDs[l] < userIDs[m] })
			up = userIDs[i]
			userIDs[i] = up

		}
	}
	return userIDs[:lenght]
}


*/

func UniqueSortedUserIDs(userIDs []int64) []int64 {
	if len(userIDs) < 2 {
		return userIDs
	}

	sort.SliceStable(userIDs, func(i, j int) bool { return userIDs[i] < userIDs[j] })
	uniqPointer := 0
	for i := 1; i < len(userIDs); i++ {
		if userIDs[uniqPointer] != userIDs[i] {
			uniqPointer++
			userIDs[uniqPointer] = userIDs[i]
		}
	}

	return userIDs[:uniqPointer+1]
}

func main() {
	l := 0
	fmt.Print("Длина слайса: ")
	fmt.Scanln(&l)
	sl := make([]int64, 0, l)

	for i := l - 1; i > -1; i-- {
		m := int64(i)
		if m%2 == 0 {
			sl = append(sl, m+1)
		} else {
			sl = append(sl, m+1)
		}
	}

	fmt.Println("Полученный слайс:", sl)
	r := UniqueSortedUserIDs(sl)
	fmt.Println("Измененный слайс:", r)
}
