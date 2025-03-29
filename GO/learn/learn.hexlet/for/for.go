package main

import (
	//"strings"
	"fmt"
)

func Map(strs []string, mapFunc func(s string) string) []string {
	fmt.Println("hui na")
	return strs
}

func mapFunc(s string) string {
	return s
}

/*
// Map iterates through the strs slice and modifies each element via mapFunc.
The func is safe and strs won't be modified.

func Map(strs []string, mapFunc func(s string) string) []string {
	mapped := make([]string, len(strs))
	for i, s := range strs {
		mapped[i] = mapFunc(s)
	}
	return mapped
}
*/

func main() {
	v := 0
	//im := 0
	//fmt.Println("Введите количество эллементов в массиве: ")
	//fmt.Scanln(&v)
	num := make([]string, 0, v)
	for i := 0; i < v; i++ {
		num = append(num, " ")
	}
	//fmt.Println("Поллученный массив: ", num)
	//fmt.Println("Ввеедите номер, который хотите удалить (от 0 до", v, "): ")
	//fmt.Scanln(&im)
	//r := Map(num, Я ХУЙ ЗНАЕТ ЧТО СЮДА ПИСАТЬ)
	// Map(strs []string, mapFunc func(s string) string) []string
	//fmt.Println(r)
}

/*
Реализуйте функцию func Map(strs []string, mapFunc func(s string) string) []string,
которая преобразует каждый элемент слайса strs с помощью функции mapFunc и возвращает новый слайс.
Учтите, что исходный слайс, который передается как strs, не должен измениться в процессе выполнения.
*/
