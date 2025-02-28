package main

import (
	"fmt"
)

func IntsCopy(src []int, maxLen int) []int {
	if maxLen <= 0 {
		fmt.Println("Такого не может быть!")
		return []int{}
	}
	if maxLen > len(src) {
		maxLen = len(src)
	}
	nums := make([]int, maxLen)
	copy(nums, src)
	return nums
}

func main() {
	l, m := 0, 0
	fmt.Println("Длина слайса:")
	fmt.Scanln(&l)
	sl := make([]int, 0, l)
	for i := 0; i < l; i++ {
		sl = append(sl, i+1)
	}
	fmt.Println("Полученный слайс:", sl)
	fmt.Println("Длина копии:")
	fmt.Scanln(&m)
	r := IntsCopy(sl, m)
	fmt.Println(r)
}

/*
func Remove(nums []int, i int) []int {
	if i <= len(nums) && i > 0 {
		l := len(nums)
		for f := i; f < l; f++ {
			nums[f-1] = nums[f]
		}
		m := l - 1
		n := make([]int, 0, m)
		for i := 0; i < m; i++ {
			n = append(n, i+1)
		}
		for l := 0; l < m; l++ {
			n[l] = nums[l]
		}
		return n
	}
	fmt.Println("Выбранный вами элемент не принадлежит массиву!")
	return nums
}

func Remove(nums []int, i int) []int {
	if i < 0 || i > len(nums)-1 {
		return nums
	}
	nums[i] = nums[len(nums)-1]
	return nums[:len(nums)-1]
}

func main() {
	v := 0
	im := 0
	fmt.Println("Введите количество эллементов в массиве: ")
	fmt.Scanln(&v)
	num := make([]int, 0, v) // первое длина, второе вместимость (будет сделан массив, куда вместится
	for i := 0; i < v; i++ { // v эллементов, но их там не будет)
		num = append(num, i+1)
	}
	fmt.Println("Поллученный массив: ", num)
	fmt.Println("Ввеедите номер, который хотите удалить (от 0 до", v, "): ")
	fmt.Scanln(&im)
	r := Remove(num, im-1)
	fmt.Println(r)
}

*/
