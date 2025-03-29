package main

import "fmt"

// func sumOfInt64(input []int64) int64 {
// 	var result int64

// 	for _, number := range input {
// 		result += number
// 	}
// 	return result
// }

// func sumOfFloat64(input []float64) float64 {
// 	var result float64

// 	for _, number := range input {
// 		result += number
// 	}
// 	return result
// }

type User struct {
	email string
	name  string
}

type Number interface { // расширеная суть интерфейса
	int64 | float64
}

// вот собственно сам дженерик, который принимает только объявленные интерфейсы

func sumOf[v Number](input []v) v {
	var result v

	for _, number := range input {
		result += number
	}
	return result
}

//

// это тоже дженерик но на основе COMPARABLE
func searchElement[C comparable](elements []C, searchEl C) bool {
	for _, el := range elements {
		if el == searchEl {
			return true
		}
	}

	return false
}

func main() {
	a := []int64{1, 2, 3}
	b := []float64{1.1, 2.2, 3.3}
	c := []string{"1", "2", "3"}

	fmt.Println(sumOf(a))
	fmt.Println(sumOf(b))

	fmt.Println(searchElement(c, "2"))

	d := []User{
		{
			email: "test",
			name:  "Nikita",
		},
		{
			email: "test2",
			name:  "Nikitos",
		},
		{
			email: "test3",
			name:  "Serega",
		},
	}

	fmt.Println(searchElement(d, User{
		email: "test3",
		name:  "Serega",
	}))

	printAny(d)
}

func printAny[A any](input A) {
	fmt.Println(input)
}
