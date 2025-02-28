package main

import "fmt"

func selectFn(n int) func(int, int) int {
	if n == 1 {
		return func(x int, y int) int { return x + y }
	} else if n == 2 {
		return func(x int, y int) int { return x - y }
	} else {
		return func(x int, y int) int { return x * y }
	}
}

func square() func() int {
	var x int = 2
	return func() int {
		x++
		return x * x
	}
}

func main() {
	m := square()
	fmt.Println(m()) // 9
	fmt.Println(m()) // 16
	fmt.Println(m()) // 25

	f := selectFn(1)
	fmt.Println(f(2, 3)) // 5
	fmt.Println(f(4, 5)) // 9
	fmt.Println(f(7, 6)) // 13
}
